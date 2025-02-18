package main

// Allow user authentication by specifying a login and a password, to gain access to the set of functionalities corresponding to this authorization level.
func (client *GbxClient) Authenticate(login string, password string) error {
	_, err := client.Call("Authenticate", login, password)
	return err
}

// Change the password for the specified login/user. Only available to SuperAdmin.
func (client *GbxClient) ChangeAuthPassword(login string, password string) error {
	_, err := client.Call("ChangeAuthPassword", login, password)
	return err
}

// Define the wanted api.
func (client *GbxClient) SetApiVersion(version string) error {
	_, err := client.Call("SetApiVersion", version)
	return err
}

// Allow the GameServer to call you back.
func (client *GbxClient) EnableCallbacks(enable bool) error {
	_, err := client.Call("EnableCallbacks", enable)
	return err
}