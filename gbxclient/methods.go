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
	ch, err := client.addCallback(id)
	if err != nil {
		return err
	}

	go client.listen()

	// Wait for connection confirmation
	select {
	case response := <-ch:
		client.deleteCallback(id)
		if response.Error != nil {
			return response.Error // Connection failed
		}
		client.Events.emit("connect", true)
		// Connection successful, return nil
		return nil
	case <-time.After(5 * time.Second):
		client.deleteCallback(id)
		client.Mutex.Lock()
		if client.Socket != nil {
			client.Socket.Close()
		}
		client.IsConnected = false
		client.Mutex.Unlock()
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
	client.Mutex.Lock()
	defer client.Mutex.Unlock()
	
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
