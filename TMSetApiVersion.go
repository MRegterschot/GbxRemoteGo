package main

import "errors"

// Define the wanted api.
func (clientt *GbxClient) SetApiVersion(version string) (bool, error) {
	res, err := clientt.Call("SetApiVersion", version)
	if err != nil {
		return false, err
	}

	// Ensure the response is a bool
	enabled, ok := res.(bool)
	if !ok {
		return false, errors.New("unexpected response format")
	}

	return enabled, nil
}