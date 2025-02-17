package main

// Display a manialink page on the client with the specified login. The first parameter is the login of the player, the other are identical to 'SendDisplayManialinkPage'. Login can be a single login or a list of comma-separated logins. Only available to Admin.
func (client *GbxClient) SendDisplayManialinkPageToLogin(login string, xml CData, timeout int, hideOnClick bool) error {
	_, err := client.Send("SendDisplayManialinkPageToLogin", login, xml, timeout, hideOnClick)
	return err
}
