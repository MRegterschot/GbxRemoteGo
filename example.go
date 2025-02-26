package main

import (
	"fmt"
	"os"

	"github.com/MRegterschot/GbxRemoteGo/events"
	. "github.com/MRegterschot/GbxRemoteGo/gbxclient"
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

	// Register gbx callback handlers
	client.OnPlayerConnect = append(client.OnPlayerConnect, func(client *GbxClient, args events.PlayerConnectEventArgs) {
		fmt.Println("Player connected:", args.Login)
	})

	client.OnPlayerCheckpoint = append(client.OnPlayerCheckpoint, func(client *GbxClient, args events.PlayerWayPointEventArgs) {
		fmt.Println("Player checkpoint:", args)
	})

	client.OnAnyCallback = append(client.OnAnyCallback, func(client *GbxClient, args CallbackEventArgs) {
		fmt.Println("Any callback:", args)
	})

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