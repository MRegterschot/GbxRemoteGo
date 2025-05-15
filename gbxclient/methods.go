package gbxclient

import (
	"errors"
	"fmt"
	"net"
	"time"
)

func (client *GbxClient) Connect() error {
	var err error

	dialTimeout := 5 * time.Second
	client.Socket, err = net.DialTimeout("tcp", net.JoinHostPort(client.Host, fmt.Sprintf("%d", client.Port)), dialTimeout)
	if err != nil {
		return err
	}
	if tcpConn, ok := client.Socket.(*net.TCPConn); ok {
		tcpConn.SetKeepAlive(true)
		tcpConn.SetKeepAlivePeriod(30 * time.Second)
	}

	id := uint32(0)
	if err := client.addCallback(id); err != nil {
		return err
	}

	go client.listen()

	// Wait for connection confirmation from handleData()
	select {
	case response := <-client.PromiseCallbacks[id]:
		delete(client.PromiseCallbacks, id) // Clean up callback
		if response.Error != nil {
			return response.Error // Connection failed
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

func (client *GbxClient) Call(method string, params ...any) (any, error) {
	if !client.IsConnected {
		return nil, errors.New("not connected")
	}

	xmlString, err := xmlSerializer(method, params)
	if err != nil {
		return nil, err
	}

	res := client.sendRequest(xmlString, true)
	return res.Result, res.Error
}

func (client *GbxClient) Send(method string, params ...any) (any, error) {
	if !client.IsConnected {
		return nil, errors.New("not connected")
	}

	xmlString, err := xmlSerializer(method, params)
	if err != nil {
		return nil, err
	}

	res := client.sendRequest(xmlString, false)
	return res.Result, res.Error
}

func (client *GbxClient) Disconnect() error {
	if client.Socket != nil {
		client.Socket.Close()
	}
	client.IsConnected = false
	return nil
}

func (e *EventEmitter) On(event string, ch chan any) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.events[event] = append(e.events[event], ch)
}

func (client *GbxClient) AddScriptCallback(method string, key string, callback func(any)) {
	if _, exists := client.ScriptCallbacks[method]; !exists {
		client.ScriptCallbacks[method] = []GbxCallbackStruct[any]{}
	}

	client.ScriptCallbacks[method] = append(client.ScriptCallbacks[method], GbxCallbackStruct[any]{
		Key:  key,
		Call: callback,
	})
}

func (client *GbxClient) RemoveScriptCallback(method string, key string) {
	if _, exists := client.ScriptCallbacks[method]; !exists {
		return
	}

	for i, cb := range client.ScriptCallbacks[method] {
		if cb.Key == key {
			client.ScriptCallbacks[method] = append(client.ScriptCallbacks[method][:i], client.ScriptCallbacks[method][i+1:]...)
			return
		}
	}
}
