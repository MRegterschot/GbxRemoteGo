package main

import "errors"

// Change the password for the specified login/user. Only available to SuperAdmin.
func (client *GbxClient) ChangeAuthPassword(login string, password string) (bool, error) {
	res, err := client.Call("ChangeAuthPassword", login, password)
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