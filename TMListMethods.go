package main

import (
	"errors"
)

// Return an array of all available XML-RPC methods on this server.
func (client *GbxClient) ListMethods() ([]string, error) {
	res, err := client.Call("system.listMethods")
	if err != nil {
		return nil, err
	}

	// Ensure the response is a slice
	data, ok := res.([]interface{})
	if !ok {
		return nil, errors.New("unexpected response format")
	}

	// Convert slice to []string
	methods := make([]string, len(data))
	for i, v := range data {
		methods[i] = v.(string)
	}

	return methods, nil
}