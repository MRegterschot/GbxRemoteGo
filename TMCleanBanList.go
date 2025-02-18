package main

// Clean the ban list of the server. Only available to Admin.
func (client *GbxClient) CleanBanList() error {
	_, err := client.Call("CleanBanList")
	return err
}