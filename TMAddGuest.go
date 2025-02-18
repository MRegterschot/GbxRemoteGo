package main

// Add the player with the specified login on the guest list. Only available to Admin.
func (client *GbxClient) AddGuest(login string) error {
	_, err := client.Call("AddGuest", login)
	return err
}