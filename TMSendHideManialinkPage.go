package main

// Hide the displayed manialink page on all clients. Only available to Admin.
func (client *GbxClient) SendHideManialinkPage() error {
	_, err := client.Call("SendHideManialinkPage")
	return err
}