package main

// Kick the player with the specified login, with an optional message. Only available to Admin.
func (client *GbxClient) Kick(login string, reason string) error {
	_, err := client.Call("Kick", login, reason)
	return err
}

// Kick the player with the specified PlayerId, with an optional message. Only available to Admin.
func (client *GbxClient) KickId(id int, reason string) error {
	_, err := client.Call("KickId", id, reason)
	return err
}