package main

import (
	"errors"
	"fmt"
	"net"
	"time"
)

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
		if err, ok := response.Error.(error); ok {
			return err // Connection failed
		}
		client.Events.emit("connect", true)
		// Connection successful, return nil
		return nil
	case <-time.After(5 * time.Second): // Timeout after 5 seconds
		delete(client.PromiseCallbacks, id) // Clean up callback
		client.Socket.Close()
		client.IsConnected = false
		client.Events.emit("disconnect", "connection timeout")
		return errors.New("connection timeout")
	}
}

func (client *GbxClient) Call(method string, params ...interface{}) (interface{}, error) {
	if !client.IsConnected {
		return nil, errors.New("not connected")
	}

	xmlString, err := xmlSerializer(method, params)
	if err != nil {
		return nil, err
	}
	
	res := client.sendRequest(xmlString)
	return res.Result, res.Error
}

func (client *GbxClient) Disconnect() error {
	if client.Socket != nil {
		client.Socket.Close()
	}
	client.IsConnected = false
	return nil
}

func (e *EventEmitter) On(event string, ch chan interface{}) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.events[event] = append(e.events[event], ch)
}
