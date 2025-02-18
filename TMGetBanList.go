package main

type BanList struct {
	Login      string `json:"Login"`
	ClientName string `json:"ClientName"`
	IPAddress  string `json:"IPAddress"`
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