package main

// Allow user authentication by specifying a login and a password, to gain access to the set of functionalities corresponding to this authorization level.
func (client *GbxClient) Authenticate(login string, password string) error {
	_, err := client.Call("Authenticate", login, password)
	return err
}