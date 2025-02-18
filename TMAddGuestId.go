package main

// Add the player with the specified PlayerId on the guest list. Only available to Admin.
func (client *GbxClient) AddGuestId(id int) error {
	_, err := client.Call("AddGuestId", id)
	return err
}