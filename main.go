package main

import (
	"fmt"
	"os"
)

func main() {
	// Create a new GbxClient
	client := NewGbxClient(Options{})

	// Register event handlers
	onConnectionChan := make(chan interface{})
	client.Events.On("connect", onConnectionChan)
	go handleConnect(onConnectionChan)

	onDisconnectChan := make(chan interface{})
	client.Events.On("disconnect", onDisconnectChan)
	go handleDisconnect(onDisconnectChan)

	onCallbackChan := make(chan interface{})
	client.Events.On("callback", onCallbackChan)
	go handleCallback(onCallbackChan)

	// Connect to the server
	if err := client.Connect("127.0.0.1", 5000); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := client.SetApiVersion("2023-04-24"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := client.EnableCallbacks(true); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := client.Authenticate("SuperAdmin", "SuperAdmin"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// if err := client.JumpToMapIndex(2); err != nil {
	// 	fmt.Println(err)
	// }

	res, err := client.AddMapList([]string{"My Maps/Beryllium.Map.Gbx"})
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
				fmt.Println("Callback received:", callback.Res)
			} else {
				fmt.Println("Invalid event type for callback.")
			}
		}
	}
}
