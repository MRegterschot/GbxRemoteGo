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
}

type XMLParam struct {
	Value string `xml:"value"`
}

type XMLRequest struct {
	XMLName    xml.Name   `xml:"methodCall"`
	MethodName string     `xml:"methodName"`
	Params     []XMLParam `xml:"params>param"`
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
	}
}

func (client *GbxClient) Connect(host string, port int) error {
	var err error
	client.Host = host
	client.Port = port
	client.Socket, err = net.Dial("tcp", fmt.Sprintf("%s:%d", client.Host, client.Port))
	if err != nil {
		return err
	}
	client.IsConnected = true
	go client.listen()
	return nil
}

func (client *GbxClient) listen() {
	buffer := make([]byte, 4096)
	for client.IsConnected {
		n, err := client.Socket.Read(buffer)
		fmt.Println("buff", string(buffer))
		if err != nil {
			if err == io.EOF {
				client.IsConnected = false
				return
			}
			continue
		}
		client.handleData(buffer[:n])
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
				}
			} else {
				g.Socket.Close()
				g.IsConnected = false
				g.Socket = nil
				// Reject connection
				if ch, ok := g.PromiseCallbacks[0]; ok {
					ch <- errors.New("GBXRemote 2 protocol not supported")
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
					ch <- []interface{}{res, err}
				}
			} else {
				// Handle method call
				method, res, err := deserializer.DeserializeMethodCall(readable)
				if err != nil {
					fmt.Println("Error:", err)
				} else {
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
func (client *GbxClient) Call(method string, params ...string) (interface{}, error) {
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

	requestData, err := xml.Marshal(xmlString)
	if err != nil {
		return nil, err
	}

	client.Mutex.Lock()
	client.RequestHandle++
	handle := client.RequestHandle
	client.Mutex.Unlock()

	ch := make(chan interface{}, 1)
	client.PromiseCallbacks[handle] = ch

	_, err = client.Socket.Write(requestData)
	if err != nil {
		delete(client.PromiseCallbacks, handle)
		return nil, err
	}

	res := <-ch
	return res, nil
}

func xmlSerializer(method string, params []string) (string, error) {
	var xmlParams []XMLParam
	for _, param := range params {
		xmlParams = append(xmlParams, XMLParam{Value: param})
	}

	request := XMLRequest{
		MethodName: method,
		Params:     xmlParams,
	}

	requestData, err := xml.Marshal(request)
	if err != nil {
		return "", err
	}

	return string(requestData), nil
}

func (client *GbxClient) Disconnect() error {
	if client.Socket != nil {
		client.Socket.Close()
	}
	client.IsConnected = false
	return nil
}

func main() {
	client := NewGbxClient(Options{})
	err := client.Connect("127.0.0.1", 5000)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Connected")

	res, err := client.Call("GetAllApiVersions")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("res", res)

	err = client.Disconnect()
	if err != nil {
		fmt.Println(err)
	}
}
