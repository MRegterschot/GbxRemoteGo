package main

type BlackList struct {
	Login      string `json:"Login"`
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