package main

import (
	"bytes"
	"encoding/base64"
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
	DataPointer      int
	Options          Options
	PromiseCallbacks map[uint32]chan interface{}
	Events           EventEmitter
}

type EventEmitter struct {
	events map[string][]chan interface{}
	mu     sync.Mutex
}

type MethodCall struct {
	MethodName string  `xml:"methodName"`
	Params     []Param `xml:"params>param"`
}
type MethodResponse struct {
	Params []Param `xml:"params>param"`
}

type Param struct {
	Value Value `xml:"value"`
}

type Value struct {
	String   string  `xml:"string"`
	Int      int     `xml:"i4"`
	Bool     bool    `xml:"boolean"`
	Float    float64 `xml:"double"`
	Struct   *Struct `xml:"struct"`
	Array    *Array  `xml:"array"`
	Base64   string  `xml:"base64"`
	DateTime string  `xml:"dateTime.iso8601"`
}

type Struct struct {
	Members []Member `xml:"member"`
}

type Member struct {
	Name  string `xml:"name"`
	Value Value  `xml:"value"`
}

type Array struct {
	Data []Value `xml:"data>value"`
}

type XMLParam struct {
	Value string `xml:",innerxml"`
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
	var methodResponse MethodResponse
	decoder := xml.NewDecoder(r)

	// Parse the XML into the methodResponse structure
	err := decoder.Decode(&methodResponse)
	if err != nil {
		return nil, err
	}

	// Extract the struct members from the XML response
	if len(methodResponse.Params) == 0 {
		return nil, errors.New("no parameters found")
	}

	param := methodResponse.Params[0]
	if param.Value.Struct == nil {
		return nil, errors.New("no struct found in value")
	}

	// Map to store the parsed data
	parsedData := make(map[string]interface{})

	// Iterate through each member of the struct and populate the map
	for _, member := range param.Value.Struct.Members {
		// Handle each type of value accordingly
		switch {
		case member.Value.String != "":
			parsedData[member.Name] = member.Value.String
		case member.Value.Int != 0:
			parsedData[member.Name] = member.Value.Int
		case member.Value.Bool:
			parsedData[member.Name] = member.Value.Bool
		case member.Value.Float != 0:
			parsedData[member.Name] = member.Value.Float
		case member.Value.Struct != nil:
			parsedData[member.Name] = member.Value.Struct
		case member.Value.Array != nil:
			parsedData[member.Name] = member.Value.Array
		case member.Value.Base64 != "":
			parsedData[member.Name] = member.Value.Base64
		case member.Value.DateTime != "":
			parsedData[member.Name] = member.Value.DateTime
		}
	}

	return parsedData, nil
}

func (d *Deserializer) DeserializeMethodCall(r io.Reader) (string, interface{}, error) {
	var methodCall MethodCall
	decoder := xml.NewDecoder(r)

	// Parse the XML into the methodCall structure
	err := decoder.Decode(&methodCall)
	if err != nil {
		return "", nil, err
	}

	// If there are no parameters, return the method name and nil for params
	if len(methodCall.Params) == 0 {
		return methodCall.MethodName, nil, nil
	}

	fmt.Println(methodCall.Params, len(methodCall.Params))

	// Parse the parameters
	params := make([]interface{}, len(methodCall.Params))
	for i, param := range methodCall.Params {
		// Here we should process different types of params (string, int, etc.)
		// For simplicity, let's assume we handle the string and integer types only
		if param.Value.String != "" {
			params[i] = param.Value.String
		} else if param.Value.Int != 0 {
			params[i] = param.Value.Int
		} else if param.Value.Bool {
			params[i] = param.Value.Bool
		} else if param.Value.Float != 0 {
			params[i] = param.Value.Float
		} else if param.Value.Array != nil {
			// Handle array if necessary
			params[i] = param.Value.Array
		} else if param.Value.Base64 != "" {
			// Handle base64 if necessary
			params[i] = param.Value.Base64
		} else if param.Value.DateTime != "" {
			// Handle dateTime if necessary
			params[i] = param.Value.DateTime
		} else if param.Value.Struct != nil {
			// Handle struct if necessary
			params[i] = param.Value.Struct
		} else {
			return "", nil, errors.New("unsupported parameter type")
		}
	}

	fmt.Println("Params:", params)

	return methodCall.MethodName, params, nil
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
		ReqHandle:        0x80000000,
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
			readable := bytes.NewReader(data[4:])

			if requestHandle >= 0x80000000 {
				// Handle method response
				// fmt.Println(string(data[4:]))
				res, err := deserializer.DeserializeMethodResponse(readable)
				if ch, ok := g.PromiseCallbacks[requestHandle]; ok {
					ch <- []interface{}{res, err}
					delete(g.PromiseCallbacks, requestHandle)
				} else {
					return
				}
			} else {
				// Handle method call
				fmt.Println(string(data[4:]))
				method, res, err := deserializer.DeserializeMethodCall(readable)
				if err != nil {
					fmt.Println("Error:", err)
				} else {
					fmt.Println("Method:", method, res)
					g.Events.Emit("callback", Callback{Method: method, Res: res})
					fmt.Println("Callback:", method, res)
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

func (client *GbxClient) Call(method string, params ...interface{}) (interface{}, error) {
	if !client.IsConnected {
		return nil, errors.New("not connected")
	}

	xmlString, err := xmlSerializer(method, params)
	if err != nil {
		return nil, err
	}

	return client.sendRequest(xmlString)
}

func (client *GbxClient) sendRequest(xmlString string) (interface{}, error) {
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

	if err := client.addCallback(handle); err != nil {
		client.Mutex.Unlock()
		return nil, err
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
		return nil, err
	}

	ch := client.PromiseCallbacks[handle]
	res := <-ch
	client.Mutex.Lock()
	delete(client.PromiseCallbacks, handle)
	client.Mutex.Unlock()

	return res, nil
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
		return fmt.Sprintf("<value><string>%s</string></value>", v), nil
	case int, int32, int64:
		return fmt.Sprintf("<value><int>%d</int></value>", v), nil
	case float32, float64:
		return fmt.Sprintf("<value><double>%f</double></value>", v), nil
	case bool:
		if v {
			return "<value><boolean>1</boolean></value>", nil
		}
		return "<value><boolean>0</boolean></value>", nil
	case []interface{}: // Handle arrays (slice of values)
		var values string
		for _, elem := range v {
			serializedElem, err := serializeParam(elem)
			if err != nil {
				return "", err
			}
			values += fmt.Sprintf("<value>%s</value>", serializedElem)
		}
		return fmt.Sprintf("<array><data>%s</data></array>", values), nil
	case []byte: // Handle base64 encoding
		encoded := base64.StdEncoding.EncodeToString(v)
		return fmt.Sprintf("<value><base64>%s</base64></value>", encoded), nil
	case time.Time: // Handle date/time serialization
		return fmt.Sprintf("<value><dateTime.iso8601>%s</dateTime.iso8601></value>", v.Format("20060102T15:04:05Z")), nil
	case map[string]interface{}: // Handle struct serialization (map of name-value pairs)
		var members string
		for name, value := range v {
			serializedValue, err := serializeParam(value)
			if err != nil {
				return "", err
			}
			members += fmt.Sprintf("<member><name>%s</name><value>%s</value></member>", name, serializedValue)
		}
		return fmt.Sprintf("<struct>%s</struct>", members), nil
	case nil: // Handle nil serialization
		return "<nil/>", nil
	default:
		// Handle unsupported types explicitly
		return "", fmt.Errorf("unsupported parameter type: %T", param)
	}
}

func (client *GbxClient) Disconnect() error {
	if client.Socket != nil {
		client.Socket.Close()
	}
	client.IsConnected = false
	return nil
}
