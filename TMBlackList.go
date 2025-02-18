package main

type BlackList struct {
	Login string `json:"Login"`
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

// UnBlackList the player with the specified login. Only available to SuperAdmin.
func (client *GbxClient) UnBlackList(login string) error {
	_, err := client.Call("UnBlackList", login)
	return err
}

// Returns the list of blacklisted players. This method takes two parameters. The first parameter specifies the maximum number of infos to be returned, and the second one the starting index in the list. The list is an array of structures. Each structure contains the following fields : Login.
func (client *GbxClient) GetBlackList(max int, start int) ([]BlackList, error) {
	res, err := client.Call("GetBlackList", max, start)
	if err != nil {
		return []BlackList{}, err
	}

	var blackList []BlackList
	err = convertToStruct(res, &blackList)
	if err != nil {
		return []BlackList{}, err
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