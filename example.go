package main

import (
	"fmt"
	"os"

	"github.com/MRegterschot/GbxRemoteGo/events"
	. "github.com/MRegterschot/GbxRemoteGo/gbxclient"
)

func main() {
	// Create a new GbxClient
	client := NewGbxClient("127.0.0.1", 5000, Options{})

	// Register event handlers
	onConnectionChan := make(chan any)
	client.Events.On("connect", onConnectionChan)
	go handleConnect(onConnectionChan)

	onDisconnectChan := make(chan any)
	client.Events.On("disconnect", onDisconnectChan)
	go handleDisconnect(onDisconnectChan)

	// Connect to the server
	if err := client.Connect(); err != nil {
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

	// Register gbx callback handlers
	client.OnPlayerConnect = append(client.OnPlayerConnect, GbxCallbackStruct[events.PlayerConnectEventArgs]{
		Key: "1",
		Call: func(args events.PlayerConnectEventArgs) {
			fmt.Println("Player connected:", args.Login)
		}})

	client.OnPlayerCheckpoint = append(client.OnPlayerCheckpoint, GbxCallbackStruct[events.PlayerWayPointEventArgs]{
		Key: "2",
		Call: func(args events.PlayerWayPointEventArgs) {
			fmt.Println("Player checkpoint:", args)
		}})

	client.OnAnyCallback = append(client.OnAnyCallback, GbxCallbackStruct[CallbackEventArgs]{
		Key: "3",
		Call: func(args CallbackEventArgs) {
			fmt.Println("Any callback:", args)
		}})

	select {}
}

func handleConnect(eventChan chan any) {
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

func handleDisconnect(eventChan chan any) {
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
