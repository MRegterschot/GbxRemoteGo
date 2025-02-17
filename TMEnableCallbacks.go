package main

import "errors"

// Allow the GameServer to call you back.
func (client *GbxClient) EnableCallbacks(enable bool) (bool, error) {
	res, err := client.Call("EnableCallbacks", enable)
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