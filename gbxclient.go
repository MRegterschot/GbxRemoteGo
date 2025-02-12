package main

import (
	"bytes"
	"encoding/binary"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Options struct {
	ShowErrors  bool
	ThrowErrors bool
}

type GbxClient struct {
	Host             string
	Port             int
	IsConnected      bool
	DoHandShake      bool
	ReqHandle        uint32
	Socket           net.Conn
	Mutex            sync.Mutex
	RecvData         bytes.Buffer
	ResponseLength   *uint32
	RequestHandle    uint32
	DataPointer      int
	Options          Options
	PromiseCallbacks map[uint32]chan interface{}
	Events           EventEmitter
}

type EventEmitter struct {
	events map[string][]chan interface{}
	mu     sync.Mutex
}

type XMLParam struct {
	Value string `xml:"value"`
}

type XMLRequest struct {
	XMLName    xml.Name   `xml:"methodCall"`
	MethodName string     `xml:"methodName"`
	Params     []XMLParam `xml:"params>param"`
}

type Callback struct {
	Method string
	Res    interface{}
}

// Deserializer placeholder (you need to implement the actual logic)
type Deserializer struct{}

func (d *Deserializer) DeserializeMethodResponse(r io.Reader) (interface{}, error) {
	// Implement XML-RPC or GBXRemote response deserialization logic here

	return nil, nil
}

func (d *Deserializer) DeserializeMethodCall(r io.Reader) (string, interface{}, error) {
	// Implement XML-RPC or GBXRemote method call deserialization logic here
	return "", nil, nil
}

func (e *EventEmitter) On(event string, ch chan interface{}) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.events[event] = append(e.events[event], ch)
}

func (e *EventEmitter) Emit(event string, value interface{}) {
	e.mu.Lock()
	defer e.mu.Unlock()
	for _, ch := range e.events[event] {
		go func(c chan interface{}) {
			c <- value
		}(ch)
	}
}

func NewGbxClient(options Options) *GbxClient {
	return &GbxClient{
		IsConnected:      false,
		Host:             "127.0.0.1",
		Port:             5000,
		Socket:           nil,
		RecvData:         bytes.Buffer{},
		ResponseLength:   nil,
		RequestHandle:    0x80000000,
		DataPointer:      0,
		DoHandShake:      false,
		Options:          options,
		PromiseCallbacks: make(map[uint32]chan interface{}),
		Events: EventEmitter{
			events: make(map[string][]chan interface{}),
		},
	}
}

func (client *GbxClient) Connect(host string, port int) error {
	var err error
	client.Host = host
	client.Port = port

	id := uint32(0)
	if err := client.addCallback(id); err != nil {
		return err
	}

	client.Socket, err = net.Dial("tcp", fmt.Sprintf("%s:%d", client.Host, client.Port))
	if err != nil {
		return err
	}
	if tcpConn, ok := client.Socket.(*net.TCPConn); ok {
		tcpConn.SetKeepAlive(true)
		tcpConn.SetKeepAlivePeriod(30 * time.Second)
	}

	go client.listen()

	// Wait for connection confirmation from handleData()
	select {
	case response := <-client.PromiseCallbacks[id]:
		delete(client.PromiseCallbacks, id) // Clean up callback
		if err, ok := response.(error); ok {
			return err // Connection failed
		}
		client.Events.Emit("connect", true)
		// Connection successful, return nil
		return nil
	case <-time.After(5 * time.Second): // Timeout after 5 seconds
		delete(client.PromiseCallbacks, id) // Clean up callback
		client.Socket.Close()
		client.IsConnected = false
		client.Events.Emit("disconnect", "connection timeout")
		return errors.New("connection timeout")
	}
}

func (client *GbxClient) addCallback(id uint32) error {
	if _, exists := client.PromiseCallbacks[id]; exists {
		return errors.New("callback already exists")
	}

	client.PromiseCallbacks[id] = make(chan interface{}, 1)
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

		fmt.Println("Received:", string(buffer[:n]))
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
					ch <- true
					delete(g.PromiseCallbacks, 0)
				}
			} else {
				g.Socket.Close()
				g.IsConnected = false
				g.Socket = nil
				// Reject connection
				if ch, ok := g.PromiseCallbacks[0]; ok {
					ch <- errors.New("GBXRemote 2 protocol not supported")
					delete(g.PromiseCallbacks, 0)
				}
			}
		} else {
			deserializer := &Deserializer{}
			requestHandle := binary.LittleEndian.Uint32(data[:4])
			fmt.Println("response:", string(data[4:]))
			readable := bytes.NewReader(data[4:])

			if requestHandle >= 0x80000000 {
				// Handle method response
				res, err := deserializer.DeserializeMethodResponse(readable)
				if ch, ok := g.PromiseCallbacks[requestHandle]; ok {
					ch <- []interface{}{res, err}
					delete(g.PromiseCallbacks, requestHandle)
				}
			} else {
				// Handle method call
				method, res, err := deserializer.DeserializeMethodCall(readable)
				if err != nil {
					fmt.Println("Error:", err)
				} else {
					g.Events.Emit("callback", Callback{Method: method, Res: res})
					fmt.Println("Callback:", method, res)
				}
			}
		}

		g.ResponseLength = nil
		if g.RecvData.Len() > 0 {
			g.handleData(nil) // Recursively handle remaining data
		}
	}
}

func (client *GbxClient) Send(method string, params ...interface{}) (interface{}, error) {
	if !client.IsConnected {
		return nil, errors.New("not connected")
	}

	xmlString, err := xmlSerializer(method, params)
	if err != nil {
		return nil, err
	}

	return client.sendRequest(xmlString, false)
}

func (client *GbxClient) Call(method string, params ...interface{}) (interface{}, error) {
	if !client.IsConnected {
		return nil, errors.New("not connected")
	}

	xmlString, err := xmlSerializer(method, params)
	if err != nil {
		return nil, err
	}

	return client.sendRequest(xmlString, true)
}

func (client *GbxClient) sendRequest(xmlString string, wait bool) (interface{}, error) {
	// if request is more than 4mb
	if len(xmlString)+8 > 4*1024*1024 {
		return nil, errors.New("request too large")
	}

	client.Mutex.Lock()
	client.ReqHandle++
	if client.ReqHandle >= 0xffffff00 {
		client.ReqHandle = 0x80000000
	}
	handle := client.ReqHandle
	client.Mutex.Unlock()

	len := len(xmlString)
	buf := make([]byte, 8+len)

	binary.LittleEndian.PutUint32(buf[0:], uint32(len))
	// Write handle as uint32 (Little Endian) at offset 4
	binary.LittleEndian.PutUint32(buf[4:], handle)
	requestData, err := xml.Marshal(xmlString)
	if err != nil {
		return nil, err
	}
	// Copy XML string into the buffer at offset 8
	copy(buf[8:], []byte(requestData))
	
	fmt.Println(string(buf))
	
	_, err = client.Socket.Write(buf)
	if err != nil {
		delete(client.PromiseCallbacks, handle)
		return nil, err
	}
	
	if (wait) {
		ch := make(chan interface{}, 1)
		client.PromiseCallbacks[handle] = ch

		res := <-ch
		fmt.Println(res)
		return res, nil
	}

	return nil, nil
}

func xmlSerializer(method string, params []interface{}) (string, error) {
	var xmlParams []XMLParam
	for _, param := range params {
		// Use reflection to handle different types of params
		paramStr, err := serializeParam(param)
		if err != nil {
			return "", err
		}
		xmlParams = append(xmlParams, XMLParam{Value: paramStr})
	}

	request := XMLRequest{
		MethodName: method,
		Params:     xmlParams,
	}

	requestData, err := xml.Marshal(request)
	if err != nil {
		return "", err
	}

	return `<?xml version="1.0"?>` + string(requestData), nil
}

// Helper function to serialize different types
func serializeParam(param interface{}) (string, error) {
	switch v := param.(type) {
	case string:
		return v, nil
	case int, int32, int64:
		return fmt.Sprintf("%d", v), nil
	case float32, float64:
		return fmt.Sprintf("%f", v), nil
	case bool:
		return fmt.Sprintf("%t", v), nil
	default:
		// Handle unknown types here, you can serialize structs or other types as needed
		return fmt.Sprintf("%v", v), nil
	}
}

func (client *GbxClient) Disconnect() error {
	if client.Socket != nil {
		client.Socket.Close()
	}
	client.IsConnected = false
	return nil
}
