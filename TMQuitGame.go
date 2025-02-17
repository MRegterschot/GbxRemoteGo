package main

import "errors"

// Quit the application. Only available to SuperAdmin.
func (client *GbxClient) QuitGame() (bool, error) {
	res, err := client.Call("QuitGame")
	if err != nil {
		return false, err
	}

	// Ensure the response is a bool
	successful, ok := res.(bool)
	if !ok {
		return false, errors.New("unexpected response format")
	}

	return successful, nil
}