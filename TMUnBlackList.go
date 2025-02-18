package main

// UnBlackList the player with the specified login. Only available to SuperAdmin.
func (client *GbxClient) UnBlackList(login string) error {
	_, err := client.Send("UnBlackList", login)
	return err
}
