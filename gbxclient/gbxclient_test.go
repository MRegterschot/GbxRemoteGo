package gbxclient

import (
	"encoding/binary"
	"net"
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

// methodResponse builds a response payload for handle: the request handle followed by the XML body.
func methodResponse(handle uint32) []byte {
	body := []byte(`<?xml version="1.0"?><methodResponse><params><param><value><boolean>1</boolean></value></param></params></methodResponse>`)
	payload := make([]byte, 4+len(body))
	binary.LittleEndian.PutUint32(payload[0:], handle)
	copy(payload[4:], body)
	return payload
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
		if _, err := client.addCallback(0); err != nil {
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
		if _, err := client.addCallback(handle); err != nil {
			t.Fatalf("addCallback: %v", err)
		}

		client.RecvData.Write(frame(methodResponse(handle)))
		assertReturns(t, "handleData", func() { client.handleData(nil) })
	})
}

// Once handleData has returned, the mutex must be free: the deadlock's second half is the caller's
// timeout path blocking in deleteCallback while the reader holds it.
func TestHandleData_ReleasesMutexForLateTimeoutCleanup(t *testing.T) {
	client := NewGbxClient("127.0.0.1", 5000, Options{})
	if _, err := client.addCallback(0); err != nil {
		t.Fatalf("addCallback: %v", err)
	}

	client.RecvData.Write(frame([]byte("GBXRemote 2")))
	assertReturns(t, "handleData", func() { client.handleData(nil) })
	assertReturns(t, "deleteCallback", func() { client.deleteCallback(0) })
}

// respondingConn hands the client the response to each request before Write returns: the order a
// server on loopback produces whenever listen() wins the race against the caller.
type respondingConn struct {
	net.Conn // only Write is used
	client   *GbxClient
}

func (c *respondingConn) Write(b []byte) (int, error) {
	handle := binary.LittleEndian.Uint32(b[4:8])
	c.client.handleData(frame(methodResponse(handle)))
	return len(b), nil
}

// A response can be handled - delivered and its callback deleted - before sendRequest is back from
// writing the request. sendRequest must still receive it, not look its channel up afterwards, find
// the entry gone and wait out its timeout on a nil channel.
func TestCall_ReceivesResponseHandledBeforeWriteReturns(t *testing.T) {
	client := NewGbxClient("127.0.0.1", 5000, Options{})
	client.IsConnected = true
	client.Socket = &respondingConn{client: client}

	if _, err := client.Call("GetVersion"); err != nil {
		t.Fatalf("Call: %v", err)
	}
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
