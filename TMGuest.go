package main

import "github.com/MRegterschot/GbxRemoteGo/structs"

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
func (client *GbxClient) GetGuestList(max int, start int) ([]structs.PlayerInfo, error) {
	res, err := client.Call("GetGuestList", max, start)
	if err != nil {
		return []structs.PlayerInfo{}, err
	}

	var guestList []structs.PlayerInfo
	err = convertToStruct(res, &guestList)
	if err != nil {
		return []structs.PlayerInfo{}, err
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
