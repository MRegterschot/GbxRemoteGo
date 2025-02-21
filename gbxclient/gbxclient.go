package gbxclient

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"reflect"
	"time"

	"github.com/MRegterschot/GbxRemoteGo/events"
	"github.com/MRegterschot/GbxRemoteGo/structs"
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
			return
		}
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
					method, parameters, err := DeserializeMethodCall(dataToProcess[4:])
					if err != nil {
						fmt.Println("Error:", err)
					} else {
						// Emit callback event
						go g.handleCallback(method, parameters)
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
	} else {
		client.Mutex.Unlock()
	}

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

	select {
	case res := <-ch:
		client.Mutex.Lock()
		delete(client.PromiseCallbacks, handle)
		client.Mutex.Unlock()
		return res
	case <-time.After(5 * time.Second):
		client.Mutex.Lock()
		delete(client.PromiseCallbacks, handle)
		client.Mutex.Unlock()
		return PromiseResult{nil, errors.New("request timed out after 5s")}
	}
}

func (client *GbxClient) handleCallback(method string, parameters []interface{}) {
	switch method {
	case "ManiaPlanet.BeginMap":
		var mapInfo structs.TMSMapInfo
		if err := convertToStruct(parameters[0], &mapInfo); err != nil {
			return
		}
		client.invokeEvents(client.OnBeginMap, events.MapEventArgs{
			Map: mapInfo,
		})
	case "ManiaPlanet.BeginMatch":
		client.invokeEventsNoArgs(client.OnBeginMatch)
	case "ManiaPlanet.Echo":
		client.invokeEvents(client.OnEcho, events.EchoEventArgs{
			Internal: parameters[0].(string),
			Public:   parameters[1].(string),
		})
	case "ManiaPlanet.EndMap":
		var mapInfo structs.TMSMapInfo
		if err := convertToStruct(parameters[0], &mapInfo); err != nil {
			return
		}
		client.invokeEvents(client.OnEndMap, events.MapEventArgs{
			Map: mapInfo,
		})
	case "ManiaPlanet.EndMatch":
		var rankings []structs.TMSPlayerRanking
		if err := convertToStruct(parameters[0], &rankings); err != nil {
			return
		}
		client.invokeEvents(client.OnEndMatch, events.EndMatchEventArgs{
			Rankings:   rankings,
			WinnerTeam: parameters[1].(int),
		})
	case "ManiaPlanet.MapListModified":
		client.invokeEvents(client.OnMapListModified, events.MapListModifiedEventArgs{
			CurMapIndex:    parameters[0].(int),
			NextMapIndex:   parameters[1].(int),
			IsListModified: parameters[2].(bool),
		})
	case "ManiaPlanet.PlayerAlliesChanged":
		client.invokeEvents(client.OnPlayerAlliesChanged, events.PlayerAlliesChangedEventArgs{
			Login: parameters[0].(string),
		})
	case "ManiaPlanet.PlayerChat":
		client.invokeEvents(client.OnPlayerChat, events.PlayerChatEventArgs{
			PlayerUid:      parameters[0].(int),
			Login:          parameters[1].(string),
			Text:           parameters[2].(string),
			IsRegistredCmd: parameters[3].(bool),
			Options:        parameters[4].(int),
		})
	case "ManiaPlanet.PlayerConnect":
		client.invokeEvents(client.OnPlayerConnect, events.PlayerConnectEventArgs{
			Login:       parameters[0].(string),
			IsSpectator: parameters[1].(bool),
		})
	case "ManiaPlanet.PlayerDisconnect":
		client.invokeEvents(client.OnPlayerDisconnect, events.PlayerDisconnectEventArgs{
			Login:  parameters[0].(string),
			Reason: parameters[1].(string),
		})
	case "ManiaPlanet.PlayerInfoChanged":
		var playerInfo structs.TMSPlayerInfo
		if err := convertToStruct(parameters[0], &playerInfo); err != nil {
			return
		}
		client.invokeEvents(client.OnPlayerInfoChanged, events.PlayerInfoChangedEventArgs{
			PlayerInfo: playerInfo,
		})
	case "ManiaPlanet.PlayerManialinkPageAnswer":
		var entries []structs.TMSEntryVal
		if err := convertToStruct(parameters[3], &entries); err != nil {
			return
		}
		client.invokeEvents(client.OnPlayerManialinkPageAnswer, events.PlayerManialinkPageAnswerEventArgs{
			PlayerUid: parameters[0].(int),
			Login:     parameters[1].(string),
			Answer:    parameters[2].(string),
			Entries:   entries,
		})
	case "ManiaPlanet.ServerStart":
		client.invokeEventsNoArgs(client.OnServerStart)
	case "ManiaPlanet.ServerStop":
		client.invokeEventsNoArgs(client.OnServerStop)
	case "ManiaPlanet.StatusChanged":
		client.invokeEvents(client.OnStatusChanged, events.StatusChangedEventArgs{
			StatusCode: parameters[0].(int),
			StatusName: parameters[1].(string),
		})
	case "ManiaPlanet.TunnelDataReceived":
		client.invokeEvents(client.OnTunnelDataReceived, events.TunnelDataReceivedEventArgs{
			PlayerUid: parameters[0].(int),
			Login:     parameters[1].(string),
			Data:      parameters[2].([]byte),
		})
	case "ManiaPlanet.VoteUpdated":
		client.invokeEvents(client.OnVoteUpdated, events.VoteUpdatedEventArgs{
			StateName: parameters[0].(string),
			Login:     parameters[1].(string),
			CmdName:   parameters[2].(string),
			CmdParam:  parameters[3].(string),
		})
	case "Trackmania.PlayerCheckpoint":
		client.invokeEvents(client.OnPlayerCheckpoint, events.PlayerCheckpointEventArgs{
			PlayerUid:       parameters[0].(int),
			Login:           parameters[1].(string),
			TimeOrScore:     parameters[2].(int),
			CurLap:          parameters[3].(int),
			CheckpointIndex: parameters[4].(int),
		})
	case "Trackmania.PlayerFinish":
		client.invokeEvents(client.OnPlayerFinish, events.PlayerFinishEventArgs{
			PlayerUid:   parameters[0].(int),
			Login:       parameters[1].(string),
			TimeOrScore: parameters[2].(int),
		})
	case "Trackmania.PlayerIncoherence":
		client.invokeEvents(client.OnPlayerIncoherence, events.PlayerIncoherenceEventArgs{
			PlayerUid: parameters[0].(int),
			Login:     parameters[1].(string),
		})
	}

	client.invokeEvents(client.OnAnyCallback, CallbackEventArgs{
		Method:     method,
		Parameters: parameters,
	})
}

// invokeEvents calls all event handlers dynamically
func (client *GbxClient) invokeEvents(events interface{}, args interface{}) {
	v := reflect.ValueOf(events)

	// Ensure `events` is a slice
	if v.Kind() != reflect.Slice {
		fmt.Println("Error: events is not a slice")
		return
	}

	for i := 0; i < v.Len(); i++ {
		handler := v.Index(i)

		// Ensure handler is a function
		if handler.Kind() != reflect.Func {
			fmt.Println("Error: event handler is not a function")
			continue
		}

		// Call the function dynamically with client and args
		handler.Call([]reflect.Value{reflect.ValueOf(client), reflect.ValueOf(args)})
	}
}

func (client *GbxClient) invokeEventsNoArgs(events interface{}) {
	v := reflect.ValueOf(events)

	// Ensure `events` is a slice
	if v.Kind() != reflect.Slice {
		fmt.Println("Error: events is not a slice")
		return
	}

	for i := 0; i < v.Len(); i++ {
		handler := v.Index(i)

		// Ensure handler is a function
		if handler.Kind() != reflect.Func {
			fmt.Println("Error: event handler is not a function")
			continue
		}

		// Call the function dynamically with client
		handler.Call([]reflect.Value{reflect.ValueOf(client)})
	}
}
