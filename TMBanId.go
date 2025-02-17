package main

// 	Ban the player with the specified PlayerId, with an optional message. Only available to Admin.
func (client *GbxClient) BanId(id int, reason string) error {
	_, err := client.Send("BanId", id, reason)
	return err
}