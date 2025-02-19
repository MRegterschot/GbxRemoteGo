package gbxclient

// Send a text message to all clients without the server login. Only available to Admin.
func (client *GbxClient) ChatSendServerMessage(message string) error {
	_, err := client.Call("ChatSendServerMessage", message)
	return err
}

// Send a text message without the server login to the client with the specified PlayerId. Only available to Admin.
func (client *GbxClient) ChatSendServerMessageToId(message string, playerId int) error {
	_, err := client.Call("ChatSendServerMessageToId", message, playerId)
	return err
}

// Send a text message without the server login to the client with the specified login. Login can be a single login or a list of comma-separated logins. Only available to Admin.
func (client *GbxClient) ChatSendServerMessageToLogin(message string, login string) error {
	_, err := client.Call("ChatSendServerMessageToLogin", message, login)
	return err
}

// Send a text message to all clients. Only available to Admin.
func (client *GbxClient) ChatSend(message string) error {
	_, err := client.Call("ChatSend", message)
	return err
}

// Send a text message to the client with the specified login. Login can be a single login or a list of comma-separated logins. Only available to Admin.
func (client *GbxClient) ChatSendToLogin(message string, login string) error {
	_, err := client.Call("ChatSendToLogin", message, login)
	return err
}

// Send a text message to the client with the specified PlayerId. Only available to Admin.
func (client *GbxClient) ChatSendToId(message string, playerId int) error {
	_, err := client.Call("ChatSendToId", message, playerId)
	return err
}

// The chat messages are no longer dispatched to the players, they only go to the rpc callback and the controller has to manually forward them. The second (optional) parameter allows all messages from the server to be automatically forwarded. Only available to Admin.
func (client *GbxClient) ChatEnableManualRouting(enable bool, forwardAll bool) error {
	_, err := client.Call("ChatEnableManualRouting", enable, forwardAll)
	return err
}

// (Text, SenderLogin, DestLogin) Send a text message to the specified DestLogin (or everybody if empty) on behalf of SenderLogin. DestLogin can be a single login or a list of comma-separated logins. Only available if manual routing is enabled. Only available to Admin.
func (client *GbxClient) ChatForwardToLogin(message string, senderLogin string, destLogin string) error {
	_, err := client.Call("ChatForwardToLogin", message, senderLogin, destLogin)
	return err
}