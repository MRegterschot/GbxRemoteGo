package main

// Ban the player with the specified login, with a message. Add it to the black list, and optionally save the new list. Only available to Admin.
func (client *GbxClient) BanAndBlackList(login string, reason string, save bool) error {
	_, err := client.Send("BanAndBlackList", login, reason, save)
	return err
}
