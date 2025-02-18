package main

// Blacklist the player with the specified PlayerId. Only available to SuperAdmin.
func (client *GbxClient) BlackListId(id int) error {
	_, err := client.Call("BlackListId", id)
	return err
}