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
	buffer := make([]byte, 8192)

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

	// Write the incoming data to RecvData buffer
	g.RecvData.Write(data)
	for g.RecvData.Len() > 4 {
		// Check if the ResponseLength is not yet set and if we have enough data to read it
		if g.ResponseLength == nil && g.RecvData.Len() >= 4 {
			lengthBytes := g.RecvData.Next(4) // Read the first 4 bytes
			length := binary.LittleEndian.Uint32(lengthBytes)

			// If connected, adjust the length by 4
			if g.IsConnected {
				length += 4
			}

			// Set the ResponseLength
			g.ResponseLength = &length
		}

		// Process the data if we have enough data for a full response
		if g.ResponseLength != nil && uint32(g.RecvData.Len()) >= *g.ResponseLength {
			// Extract the exact amount of data based on the response length
			dataToProcess := g.RecvData.Next(int(*g.ResponseLength))

			// Handle the connection and disconnection logic
			if !g.IsConnected {
				// Check if we received the connection response
				if string(dataToProcess) == "GBXRemote 2" {
					g.IsConnected = true
					// Resolve connection success
					if ch, ok := g.PromiseCallbacks[0]; ok {
						ch <- PromiseResult{Result: true, Error: nil}
						delete(g.PromiseCallbacks, 0)
					}
				} else {
					// Connection failed, close socket and reject promise
					g.Socket.Close()
					g.IsConnected = false
					g.Socket = nil
					if ch, ok := g.PromiseCallbacks[0]; ok {
						ch <- PromiseResult{Result: false, Error: errors.New("connection failed")}
						delete(g.PromiseCallbacks, 0)
					}
					g.Events.emit("disconnect", "Connection failed")
				}
			} else {
				// If connected, process method response or method call
				requestHandle := binary.LittleEndian.Uint32(dataToProcess[:4])

				if requestHandle >= 0x80000000 {
					// Handle method response
					res, err := DeserializeMethodResponse(dataToProcess[4:])
					if ch, ok := g.PromiseCallbacks[requestHandle]; ok {
						ch <- PromiseResult{Result: res, Error: err}
						delete(g.PromiseCallbacks, requestHandle)
					}
				} else {
					// Handle method call
					method, res, err := DeserializeMethodCall(dataToProcess[4:])
					if err != nil {
						fmt.Println("Error:", err)
					} else {
						// Emit callback event
						g.Events.emit("callback", Callback{Method: method, Res: res})
					}
				}
			}

			// Reset the response length after processing
			g.ResponseLength = nil
		}
	}
}

func (client *GbxClient) sendRequest(xmlString string, wait bool) PromiseResult {
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

	if wait {
		if err := client.addCallback(handle); err != nil {
			client.Mutex.Unlock()
			return PromiseResult{nil, err}
		}
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

	if !wait {
		return PromiseResult{nil, nil}
	}

	ch := client.PromiseCallbacks[handle]
	res := <-ch
	client.Mutex.Lock()
	delete(client.PromiseCallbacks, handle)
	client.Mutex.Unlock()

	return res
}
