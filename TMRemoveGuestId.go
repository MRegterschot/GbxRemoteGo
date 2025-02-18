package main

// Remove the player with the specified PlayerId from the guest list. Only available to Admin.
func (client *GbxClient) RemoveGuestId(id int) error {
	_, err := client.Call("RemoveGuestId", id)
	return err
}