package main

// Clean the guest list of the server. Only available to Admin.
func (client *GbxClient) CleanGuestList() error {
	_, err := client.Call("CleanGuestList")
	return err
}