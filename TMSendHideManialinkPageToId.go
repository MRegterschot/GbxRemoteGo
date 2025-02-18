package main

// Hide the displayed manialink page on the client with the specified UId. Only available to Admin.
func (client *GbxClient) SendHideManialinkPageToId(uid int) error {
	_, err := client.Call("SendHideManialinkPageToId", uid)
	return err
}