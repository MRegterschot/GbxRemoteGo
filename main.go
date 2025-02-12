package main

import (
	"fmt"
)

func main() {
	client := NewGbxClient(Options{})
	onConnectionChan := make(chan interface{})
	client.Events.On("connect", onConnectionChan)
	go handleConnect(onConnectionChan)

	onDisconnectChan := make(chan interface{})
	client.Events.On("disconnect", onDisconnectChan)
	go handleDisconnect(onDisconnectChan)

	onCallbackChan := make(chan interface{})
	client.Events.On("callback", onCallbackChan)
	go handleCallback(onCallbackChan)

	if err := client.Connect("127.0.0.1", 5000); err != nil {
		fmt.Println(err)
		return
	}

	client.Send("SetApiVersion", "2023-04-24")
	client.Send("EnableCallbacks", true)

	if _, err := client.Call("Authenticate", "SuperAdmin", "SuperAdmin"); err != nil {
		fmt.Println(err)
		return
	}

	res, err := client.Call("GetSystemInfo")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)

	select {}
}

func handleConnect(eventChan chan interface{}) {
	for {
		select {
		case event := <-eventChan:
			if connected, ok := event.(bool); ok {
				if connected {
					fmt.Println("Connected")
				} else {
					fmt.Println("Not Connected")
				}
			} else {
				fmt.Println("Invalid event type for connect.")
			}
		}
	}
}

// handleDisconnect processes the disconnect message
func handleDisconnect(eventChan chan interface{}) {
	for {
		select {
		case event := <-eventChan:
			if msg, ok := event.(string); ok {
				fmt.Println(msg)
			} else {
				fmt.Println("Invalid event type for disconnect.")
			}
		}
	}
}

// handleCallback processes the callback event
func handleCallback(eventChan chan interface{}) {
	for {
		select {
		case event := <-eventChan:
			if callback, ok := event.(Callback); ok {
				fmt.Println("Callback received:", callback)
			} else {
				fmt.Println("Invalid event type for callback.")
			}
		}
	}
}
