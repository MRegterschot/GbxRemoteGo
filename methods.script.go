package main

import "errors"

// Get the current mode script.
func (client *GbxClient) GetModeScriptText() (string, error) {
	res, err := client.Call("GetModeScriptText")
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

// Set the mode script and restart. Only available to Admin.
func (client *GbxClient) SetModeScriptText(script string) error {
	_, err := client.Call("SetModeScriptText", script)
	return err
}
