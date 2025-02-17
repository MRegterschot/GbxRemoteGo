package main

import "errors"

// (debug tool) Connect a fake player to the server and returns the login. Only available to Admin.
func (client *GbxClient) ConnectFakePlayer() (string, error) {
	res, err := client.Call("ConnectFakePlayer")
	if err != nil {
		return "", err
	}

	login, ok := res.(string)
	if !ok {
		return "", errors.New("unexpected response format")
	}

	return login, nil
}