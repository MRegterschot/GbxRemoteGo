package main

import "errors"

// Display a manialink page on all clients. The parameters are the xml description of the page to display, a timeout to autohide it (0 = permanent), and a boolean to indicate whether the page must be hidden as soon as the user clicks on a page option. Only available to Admin.
func (client *GbxClient) SendDisplayManialinkPage(xml CData, timeout int, hideOnClick bool) (bool, error) {
	res, err := client.Call("SendDisplayManialinkPage", xml, timeout, hideOnClick)
	if err != nil {
		return false, err
	}

	successful, ok := res.(bool)
	if !ok {
		return false, errors.New("unexpected response format")
	}

	return successful, nil
}
