package main

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

func NewGbxClient(options Options) *GbxClient {
	return &GbxClient{
		IsConnected:      false,
		Host:             "127.0.0.1",
		Port:             5000,
		Socket:           nil,
		RecvData:         bytes.Buffer{},
		ResponseLength:   nil,
		ReqHandle:        0x80000000,
		DataPointer:      0,
		DoHandShake:      false,
		Options:          options,
		PromiseCallbacks: make(map[uint32]chan PromiseResult),
		Events: EventEmitter{
			events: make(map[string][]chan interface{}),
		},
	}
}

func (e *EventEmitter) emit(event string, value interface{}) {
	e.mu.Lock()
	defer e.mu.Unlock()
	for _, ch := range e.events[event] {
		go func(c chan interface{}) {
			c <- value
		}(ch)
	}
}

func (client *GbxClient) addCallback(id uint32) error {
	if _, exists := client.PromiseCallbacks[id]; exists {
		return errors.New("callback already exists")
	}

	client.PromiseCallbacks[id] = make(chan PromiseResult)
	return nil
}

func (client *GbxClient) listen() {
	buffer := make([]byte, 4096)

	for {
		n, err := client.Socket.Read(buffer)
		if err != nil {
			if err == io.EOF {
				client.IsConnected = false
				return
			}
			fmt.Println("Read error:", err)
			return
		}
		// fmt.Println("Received data:", string(buffer[:n]))
		client.handleData(buffer[:n]) // Pass only received data
	}
}

func (g *GbxClient) handleData(data []byte) {
	g.Mutex.Lock()
	defer g.Mutex.Unlock()

	g.RecvData.Write(data)

	if g.RecvData.Len() > 0 && g.ResponseLength == nil {
		if g.RecvData.Len() < 4 {
			return // Not enough data to read the length
		}

		lengthBytes := g.RecvData.Next(4)
		length := binary.LittleEndian.Uint32(lengthBytes)
		if g.IsConnected {
			length += 4
		}
		g.ResponseLength = &length
	}

	if g.ResponseLength != nil && uint32(g.RecvData.Len()) >= *g.ResponseLength {
		data := g.RecvData.Next(int(*g.ResponseLength))

		if !g.IsConnected {
			if string(data) == "GBXRemote 2" {
				g.IsConnected = true
				// Resolve connection success
				if ch, ok := g.PromiseCallbacks[0]; ok {
					ch <- PromiseResult{Result: true, Error: nil}
					delete(g.PromiseCallbacks, 0)
				}
			} else {
				g.Socket.Close()
				g.IsConnected = false
				g.Socket = nil
				// Reject connection
				if ch, ok := g.PromiseCallbacks[0]; ok {
					ch <- PromiseResult{Result: false, Error: errors.New("connection failed")}
					delete(g.PromiseCallbacks, 0)
				}
			}
		} else {
			deserializer := &Deserializer{}
			requestHandle := binary.LittleEndian.Uint32(data[:4])
			readable := bytes.NewReader(data[4:])

			if requestHandle >= 0x80000000 {
				// Handle method response
				res, err := deserializer.DeserializeMethodResponse(readable)
				if ch, ok := g.PromiseCallbacks[requestHandle]; ok {
					ch <- PromiseResult{Result: res, Error: err}
					delete(g.PromiseCallbacks, requestHandle)
				} else {
					return
				}
			} else {
				// Handle method call
				method, res, err := deserializer.DeserializeMethodCall(readable)
				if err != nil {
					fmt.Println("Error:", err)
				} else {
					g.Events.emit("callback", Callback{Method: method, Res: res})
				}
			}
		}

		g.ResponseLength = nil
		if g.RecvData.Len() > 0 {
			g.handleData(nil) // Recursively handle remaining data
		}
		return
	}
	return
}

func (client *GbxClient) sendRequest(xmlString string) PromiseResult {
	// if request is more than 4mb
	if len(xmlString)+8 > 4*1024*1024 {
		return PromiseResult{nil, errors.New("request too large")}
	}

	client.Mutex.Lock()
	client.ReqHandle++
	if client.ReqHandle >= 0xffffff00 {
		client.ReqHandle = 0x80000000
	}
	handle := client.ReqHandle

	if err := client.addCallback(handle); err != nil {
		client.Mutex.Unlock()
		return PromiseResult{nil, err}
	}
	client.Mutex.Unlock()

	len := len(xmlString)
	buf := make([]byte, 8+len)

	binary.LittleEndian.PutUint32(buf[0:], uint32(len))
	// Write handle as uint32 (Little Endian) at offset 4
	binary.LittleEndian.PutUint32(buf[4:], handle)
	// Copy XML string into the buffer at offset 8
	copy(buf[8:], []byte(xmlString))

	_, err := client.Socket.Write(buf)
	if err != nil {
		client.Mutex.Lock()
		delete(client.PromiseCallbacks, handle)
		client.Mutex.Unlock()
		return PromiseResult{nil, err}
	}

	ch := client.PromiseCallbacks[handle]
	res := <-ch
	client.Mutex.Lock()
	delete(client.PromiseCallbacks, handle)
	client.Mutex.Unlock()

	return res
}
