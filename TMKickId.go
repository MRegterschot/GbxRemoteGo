package main

// Kick the player with the specified PlayerId, with an optional message. Only available to Admin.
func (client *GbxClient) KickId(id int, reason string) error {
	_, err := client.Send("KickId", id, reason)
	return err
}