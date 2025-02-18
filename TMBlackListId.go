package main

// Blacklist the player with the specified PlayerId. Only available to SuperAdmin.
func (client *GbxClient) BlackListId(id int) error {
	_, err := client.Send("BlackListId", id)
	return err
}