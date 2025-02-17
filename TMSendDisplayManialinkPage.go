package main

// Display a manialink page on all clients. The parameters are the xml description of the page to display, a timeout to autohide it (0 = permanent), and a boolean to indicate whether the page must be hidden as soon as the user clicks on a page option. Only available to Admin.
func (client *GbxClient) SendDisplayManialinkPage(xml CData, timeout int, hideOnClick bool) error {
	_, err := client.Send("SendDisplayManialinkPage", xml, timeout, hideOnClick)
	return err
}
