package main

import "errors"

// Display a manialink page on the client with the specified UId. The first parameter is the UId of the player, the other are identical to 'SendDisplayManialinkPage'. Only available to Admin.
func (client *GbxClient) SendDisplayManialinkPageToId(uid int, xml CData, timeout int, hideOnClick bool) (bool, error) {
	res, err := client.Call("SendDisplayManialinkPageToId", uid, xml, timeout, hideOnClick)
	if err != nil {
		return false, err
	}

	successful, ok := res.(bool)
	if !ok {
		return false, errors.New("unexpected response format")
	}

	return successful, nil
}
