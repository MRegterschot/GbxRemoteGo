package main

// Unban the player with the specified login. Only available to Admin.
func (client *GbxClient) UnBan(login string) error {
	_, err := client.Call("UnBan", login)
	return err
}
