package main

// (debug tool) Connect a fake player to the server. Only available to Admin.
func (client *GbxClient) ConnectFakePlayer() error {
	_, err := client.Send("ConnectFakePlayer")
	return err
}

// (debug tool) Disconnect a fake player, or all the fake players if login is '*'. Only available to Admin.
func (client *GbxClient) DisconnectFakePlayer(login string) error {
	_, err := client.Send("DisconnectFakePlayer", login)
	return err
}
