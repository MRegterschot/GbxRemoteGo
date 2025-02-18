package main

type PlayerInfo struct {
	Login string `xmlrpc:"Login"`
}

// Returns the list of players on the guest list. This method takes two parameters. The first parameter specifies the maximum number of infos to be returned, and the second one the starting index in the list. The list is an array of structures. Each structure contains the following fields : Login.
func (client *GbxClient) GetGuestList(max int, start int) ([]PlayerInfo, error) {
	res, err := client.Call("GetGuestList", max, start)
	if err != nil {
		return []PlayerInfo{}, err
	}

	var guestList []PlayerInfo
	err = convertToStruct(res, &guestList)
	if err != nil {
		return []PlayerInfo{}, err
	}

	return guestList, nil
}