package main

// 	Ban the player with the specified PlayerId, with an optional message. Only available to Admin.
func (client *GbxClient) BanId(id int, reason string) error {
	_, err := client.Call("BanId", id, reason)
	return err
}