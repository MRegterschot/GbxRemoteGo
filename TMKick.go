package main

// Kick the player with the specified login, with an optional message. Only available to Admin.
func (client *GbxClient) Kick(login string, reason string) error {
	_, err := client.Call("Kick", login, reason)
	return err
}