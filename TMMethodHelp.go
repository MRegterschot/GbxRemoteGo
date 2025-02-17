package main

import (
	"errors"
)

// Given the name of a method, return a help string.
func (client *GbxClient) MethodHelp(method string) (string, error) {
	res, err := client.Call("system.methodHelp", method)
	if err != nil {
		return "", err
	}

	// Ensure the response is a string
	data, ok := res.(string)
	if !ok {
		return "", errors.New("unexpected response format")
	}

	return data, nil
}