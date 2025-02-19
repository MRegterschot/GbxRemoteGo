package gbxclient

// Display a manialink page on all clients. The parameters are the xml description of the page to display, a timeout to autohide it (0 = permanent), and a boolean to indicate whether the page must be hidden as soon as the user clicks on a page option. Only available to Admin.
func (client *GbxClient) SendDisplayManialinkPage(xml CData, timeout int, hideOnClick bool) error {
	_, err := client.Call("SendDisplayManialinkPage", xml, timeout, hideOnClick)
	return err
}

// Hide the displayed manialink page on all clients. Only available to Admin.
func (client *GbxClient) SendHideManialinkPage() error {
	_, err := client.Call("SendHideManialinkPage")
	return err
}

// Hide the displayed manialink page on the client with the specified UId. Only available to Admin.
func (client *GbxClient) SendHideManialinkPageToId(uid int) error {
	_, err := client.Call("SendHideManialinkPageToId", uid)
	return err
}

// Display a manialink page on the client with the specified UId. The first parameter is the UId of the player, the other are identical to 'SendDisplayManialinkPage'. Only available to Admin.
func (client *GbxClient) SendDisplayManialinkPageToId(uid int, xml CData, timeout int, hideOnClick bool) error {
	_, err := client.Call("SendDisplayManialinkPageToId", uid, xml, timeout, hideOnClick)
	return err
}

// Hide the displayed manialink page on the client with the specified login. Login can be a single login or a list of comma-separated logins. Only available to Admin.
func (client *GbxClient) SendHideManialinkPageToLogin(login string) error {
	_, err := client.Call("SendHideManialinkPageToLogin", login)
	return err
}

// Display a manialink page on the client with the specified login. The first parameter is the login of the player, the other are identical to 'SendDisplayManialinkPage'. Login can be a single login or a list of comma-separated logins. Only available to Admin.
func (client *GbxClient) SendDisplayManialinkPageToLogin(login string, xml CData, timeout int, hideOnClick bool) error {
	_, err := client.Call("SendDisplayManialinkPageToLogin", login, xml, timeout, hideOnClick)
	return err
}
