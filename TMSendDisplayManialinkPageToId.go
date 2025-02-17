package main

// Display a manialink page on the client with the specified UId. The first parameter is the UId of the player, the other are identical to 'SendDisplayManialinkPage'. Only available to Admin.
func (client *GbxClient) SendDisplayManialinkPageToId(uid int, xml CData, timeout int, hideOnClick bool) error {
	_, err := client.Send("SendDisplayManialinkPageToId", uid, xml, timeout, hideOnClick)
	return err
}
