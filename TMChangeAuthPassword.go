package main

// Change the password for the specified login/user. Only available to SuperAdmin.
func (client *GbxClient) ChangeAuthPassword(login string, password string) error {
	_, err := client.Call("ChangeAuthPassword", login, password)
	return err
}