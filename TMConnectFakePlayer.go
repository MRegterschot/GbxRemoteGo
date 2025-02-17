package main

// (debug tool) Connect a fake player to the server. Only available to Admin.
func (client *GbxClient) ConnectFakePlayer() error {
	_, err := client.Send("ConnectFakePlayer")
	return err
}