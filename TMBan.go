package main

// Ban the player with the specified login, with an optional message. Only available to Admin.
func (client *GbxClient) Ban(login string, reason string) error {
	_, err := client.Call("Ban", login, reason)
	return err
}