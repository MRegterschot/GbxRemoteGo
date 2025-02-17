package main

type TMSystemInfo struct {
	PublishedIp             string `json:"PublishedIp"`
	Port                    int    `json:"Port"`
	P2PPort                 int    `json:"P2PPort"`
	TitleId                 string `json:"TitleId"`
	ServerLogin             string `json:"ServerLogin"`
	ServerPlayerId          int    `json:"ServerPlayerId"`
	ConnectionDownloadRate  int    `json:"ConnectionDownloadRate"`
	ConnectionUploadRate    int    `json:"ConnectionUploadRate"`
	IsServer                bool   `json:"IsServer"`
	IsDedicated             bool   `json:"IsDedicated"`
}

// Get some system infos, including connection rates (in kbps).
func (client *GbxClient) GetSystemInfo() (TMSystemInfo, error) {
	res, err := client.Call("GetSystemInfo")
	if err != nil {
		return TMSystemInfo{}, err
	}

	var systemInfo TMSystemInfo
	err = convertToStruct(res, &systemInfo)
	if err != nil {
		return TMSystemInfo{}, err
	}

	return systemInfo, nil
}