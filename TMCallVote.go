package main

import "errors"

func (client *GbxClient) CallVote(vote string) (bool, error) {
	res, err := client.Call("CallVote", vote)
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