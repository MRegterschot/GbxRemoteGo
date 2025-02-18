package main

// Clean the blacklist of the server. Only available to SuperAdmin.
func (client *GbxClient) CleanBlackList() error {
	_, err := client.Call("CleanBlackList")
	return err
}