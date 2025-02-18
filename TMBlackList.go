package main

// Blacklist the player with the specified login. Only available to SuperAdmin.
func (client *GbxClient) BlackList(login string) error {
	_, err := client.Send("BlackList", login)
	return err
}