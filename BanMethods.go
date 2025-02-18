package main

type BanList struct {
	Login      string `json:"Login"`
	ClientName string `json:"ClientName"`
	IPAddress  string `json:"IPAddress"`
}

// Ban the player with the specified login, with an optional message. Only available to Admin.
func (client *GbxClient) Ban(login string, reason string) error {
	_, err := client.Call("Ban", login, reason)
	return err
}

// Ban the player with the specified login, with a message. Add it to the black list, and optionally save the new list. Only available to Admin.
func (client *GbxClient) BanAndBlackList(login string, reason string, save bool) error {
	_, err := client.Call("BanAndBlackList", login, reason, save)
	return err
}

// 	Ban the player with the specified PlayerId, with an optional message. Only available to Admin.
func (client *GbxClient) BanId(id int, reason string) error {
	_, err := client.Call("BanId", id, reason)
	return err
}

// Unban the player with the specified login. Only available to Admin.
func (client *GbxClient) UnBan(login string) error {
	_, err := client.Call("UnBan", login)
	return err
}

// 	Returns the list of banned players. This method takes two parameters. The first parameter specifies the maximum number of infos to be returned, and the second one the starting index in the list. The list is an array of structures. Each structure contains the following fields : Login, ClientName and IPAddress.
func (client *GbxClient) GetBanList(max int, start int) ([]BanList, error) {
	res, err := client.Call("GetBanList", max, start)
	if err != nil {
		return []BanList{}, err
	}

	var banList []BanList
	err = convertToStruct(res, &banList)
	if err != nil {
		return []BanList{}, err
	}

	return banList, nil
}

// Clean the ban list of the server. Only available to Admin.
func (client *GbxClient) CleanBanList() error {
	_, err := client.Call("CleanBanList")
	return err
}