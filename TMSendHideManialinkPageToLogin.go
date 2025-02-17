package main

// Hide the displayed manialink page on the client with the specified login. Login can be a single login or a list of comma-separated logins. Only available to Admin.
func (client *GbxClient) SendHideManialinkPageToLogin(login string) error {
	_, err := client.Send("SendHideManialinkPageToLogin", login)
	return err
}