package main

import "errors"

// Cancel the current vote. Only available to Admin.
func (client *GbxClient) CancelVote() (bool, error) {
	res, err := client.Call("CancelVote")
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