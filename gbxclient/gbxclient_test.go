package gbxclient

import (
	"encoding/binary"
	"testing"
	"time"
)

// frame builds a raw server frame: a little-endian length followed by the payload, which is how
// listen() hands data to handleData.
func frame(payload []byte) []byte {
	buf := make([]byte, 4+len(payload))
	binary.LittleEndian.PutUint32(buf[0:], uint32(len(payload)))
	copy(buf[4:], payload)
	return buf
}

// handleData writes results into the promise channel while holding Mutex. If that send can block,
// a response that arrives once the caller has stopped receiving - which is exactly what
// sendRequest/Connect do when their 5s timeout fires - wedges the reader goroutine with the mutex
// held, and every later call deadlocks against it.
//
// Both cases below register a callback that nobody ever reads from, then hand handleData the
// matching response. They must return regardless.
func TestHandleData_DoesNotBlockWhenNobodyIsReceiving(t *testing.T) {
	t.Run("connect handshake", func(t *testing.T) {
		client := NewGbxClient("127.0.0.1", 5000, Options{})
		if err := client.addCallback(0); err != nil {
			t.Fatalf("addCallback: %v", err)
		}

		client.RecvData.Write(frame([]byte("GBXRemote 2")))
		assertReturns(t, "handleData", func() { client.handleData(nil) })

		if !client.IsConnected {
			t.Error("expected the handshake to mark the client connected")
		}
	})

	t.Run("method response", func(t *testing.T) {
		client := NewGbxClient("127.0.0.1", 5000, Options{})
		client.IsConnected = true

		const handle uint32 = 0x80000001
		if err := client.addCallback(handle); err != nil {
			t.Fatalf("addCallback: %v", err)
		}

		// A response frame is the request handle followed by the XML body. The body does not have
		// to deserialize cleanly - handleData forwards the error over the same channel, which is
		// the send under test.
		body := []byte(`<?xml version="1.0"?><methodResponse><params><param><value><boolean>1</boolean></value></param></params></methodResponse>`)
		payload := make([]byte, 4+len(body))
		binary.LittleEndian.PutUint32(payload[0:], handle)
		copy(payload[4:], body)

		client.RecvData.Write(frame(payload))
		assertReturns(t, "handleData", func() { client.handleData(nil) })
	})
}

// Once handleData has returned, the mutex must be free: the deadlock's second half is the caller's
// timeout path blocking in deleteCallback while the reader holds it.
func TestHandleData_ReleasesMutexForLateTimeoutCleanup(t *testing.T) {
	client := NewGbxClient("127.0.0.1", 5000, Options{})
	if err := client.addCallback(0); err != nil {
		t.Fatalf("addCallback: %v", err)
	}

	client.RecvData.Write(frame([]byte("GBXRemote 2")))
	assertReturns(t, "handleData", func() { client.handleData(nil) })
	assertReturns(t, "deleteCallback", func() { client.deleteCallback(0) })
}

// assertReturns fails the test rather than hanging it when fn deadlocks.
func assertReturns(t *testing.T, name string, fn func()) {
	t.Helper()
	done := make(chan struct{})
	go func() {
		defer close(done)
		fn()
	}()
	select {
	case <-done:
	case <-time.After(2 * time.Second):
		t.Fatalf("%s blocked: the promise channel send deadlocked while holding Mutex", name)
	}
}
