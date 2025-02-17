package main

import "errors"

func (client *GbxClient) CallVoteEx(vote string, ratio float32, timeout int, who int) (bool, error) {
	res, err := client.Call("CallVoteEx", vote, ratio, timeout, who)
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