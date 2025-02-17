package main

// (debug tool) Disconnect a fake player, or all the fake players if login is '*'. Only available to Admin.
func (client *GbxClient) DisconnectFakePlayer(login string) error {
	_, err := client.Send("DisconnectFakePlayer", login)
	return err
}