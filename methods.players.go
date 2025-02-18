package main

import "github.com/MRegterschot/GbxRemoteGo/structs"

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

// Ban the player with the specified PlayerId, with an optional message. Only available to Admin.
func (client *GbxClient) BanId(id int, reason string) error {
	_, err := client.Call("BanId", id, reason)
	return err
}

// Unban the player with the specified login. Only available to Admin.
func (client *GbxClient) UnBan(login string) error {
	_, err := client.Call("UnBan", login)
	return err
}

// Returns the list of banned players. This method takes two parameters. The first parameter specifies the maximum number of infos to be returned, and the second one the starting index in the list. The list is an array of structures. Each structure contains the following fields : Login, ClientName and IPAddress.
func (client *GbxClient) GetBanList(max int, start int) ([]structs.TMBanList, error) {
	res, err := client.Call("GetBanList", max, start)
	if err != nil {
		return []structs.TMBanList{}, err
	}

	var banList []structs.TMBanList
	err = convertToStruct(res, &banList)
	if err != nil {
		return []structs.TMBanList{}, err
	}

	return banList, nil
}

// Clean the ban list of the server. Only available to Admin.
func (client *GbxClient) CleanBanList() error {
	_, err := client.Call("CleanBanList")
	return err
}

// Blacklist the player with the specified login. Only available to SuperAdmin.
func (client *GbxClient) BlackList(login string) error {
	_, err := client.Call("BlackList", login)
	return err
}

// Blacklist the player with the specified PlayerId. Only available to SuperAdmin.
func (client *GbxClient) BlackListId(id int) error {
	_, err := client.Call("BlackListId", id)
	return err
}

// UnTMBlackList the player with the specified login. Only available to SuperAdmin.
func (client *GbxClient) UnBlackList(login string) error {
	_, err := client.Call("UnBlackList", login)
	return err
}

// Returns the list of blacklisted players. This method takes two parameters. The first parameter specifies the maximum number of infos to be returned, and the second one the starting index in the list. The list is an array of structures. Each structure contains the following fields : Login.
func (client *GbxClient) GetBlackList(max int, start int) ([]structs.TMBlackList, error) {
	res, err := client.Call("GetBlackList", max, start)
	if err != nil {
		return []structs.TMBlackList{}, err
	}

	var blackList []structs.TMBlackList
	err = convertToStruct(res, &blackList)
	if err != nil {
		return []structs.TMBlackList{}, err
	}

	return blackList, nil
}

// Load the black list file with the specified file name. Only available to Admin.
func (client *GbxClient) LoadBlackList(fileName string) error {
	_, err := client.Call("LoadBlackList", fileName)
	return err
}

// Save the black list in the file with specified file name. Only available to Admin.
func (client *GbxClient) SaveBlackList(fileName string) error {
	_, err := client.Call("SaveBlackList", fileName)
	return err
}

// Clean the blacklist of the server. Only available to SuperAdmin.
func (client *GbxClient) CleanBlackList() error {
	_, err := client.Call("CleanBlackList")
	return err
}

// (debug tool) Connect a fake player to the server. Only available to Admin.
func (client *GbxClient) ConnectFakePlayer() error {
	_, err := client.Send("ConnectFakePlayer")
	return err
}

// (debug tool) Disconnect a fake player, or all the fake players if login is '*'. Only available to Admin.
func (client *GbxClient) DisconnectFakePlayer(login string) error {
	_, err := client.Send("DisconnectFakePlayer", login)
	return err
}

// Add the player with the specified login on the guest list. Only available to Admin.
func (client *GbxClient) AddGuest(login string) error {
	_, err := client.Call("AddGuest", login)
	return err
}

// Add the player with the specified PlayerId on the guest list. Only available to Admin.
func (client *GbxClient) AddGuestId(id int) error {
	_, err := client.Call("AddGuestId", id)
	return err
}

// Remove the player with the specified login from the guest list. Only available to Admin.
func (client *GbxClient) RemoveGuest(login string) error {
	_, err := client.Call("RemoveGuest", login)
	return err
}

// Returns the list of players on the guest list. This method takes two parameters. The first parameter specifies the maximum number of infos to be returned, and the second one the starting index in the list. The list is an array of structures. Each structure contains the following fields : Login.
func (client *GbxClient) GetGuestList(max int, start int) ([]structs.TMPlayerInfo, error) {
	res, err := client.Call("GetGuestList", max, start)
	if err != nil {
		return []structs.TMPlayerInfo{}, err
	}

	var guestList []structs.TMPlayerInfo
	err = convertToStruct(res, &guestList)
	if err != nil {
		return []structs.TMPlayerInfo{}, err
	}

	return guestList, nil
}

// Load the guest list file with the specified file name. Only available to Admin.
func (client *GbxClient) LoadGuestList(fileName string) error {
	_, err := client.Call("LoadGuestList", fileName)
	return err
}

// Remove the player with the specified PlayerId from the guest list. Only available to Admin.
func (client *GbxClient) RemoveGuestId(id int) error {
	_, err := client.Call("RemoveGuestId", id)
	return err
}

// Save the guest list in the file with specified file name. Only available to Admin.
func (client *GbxClient) SaveGuestList(fileName string) error {
	_, err := client.Call("SaveGuestList", fileName)
	return err
}

// Clean the guest list of the server. Only available to Admin.
func (client *GbxClient) CleanGuestList() error {
	_, err := client.Call("CleanGuestList")
	return err
}

// Kick the player with the specified login, with an optional message. Only available to Admin.
func (client *GbxClient) Kick(login string, reason string) error {
	_, err := client.Call("Kick", login, reason)
	return err
}

// Kick the player with the specified PlayerId, with an optional message. Only available to Admin.
func (client *GbxClient) KickId(id int, reason string) error {
	_, err := client.Call("KickId", id, reason)
	return err
}
