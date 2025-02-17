package main

import "errors"

// Allow user authentication by specifying a login and a password, to gain access to the set of functionalities corresponding to this authorization level.
func (client *GbxClient) Authenticate(login string, password string) (bool, error) {
	res, err := client.Call("Authenticate", login, password)
	if err != nil {
		return false, err
	}

	// Ensure the response is a bool
	authenticated, ok := res.(bool)
	if !ok {
		return false, errors.New("unexpected response format")
	}

	return authenticated, nil
}