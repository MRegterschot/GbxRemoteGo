package main

type TMSystemInfo struct {
	PublishedIp            string `json:"PublishedIp"`
	Port                   int    `json:"Port"`
	P2PPort                int    `json:"P2PPort"`
	TitleId                string `json:"TitleId"`
	ServerLogin            string `json:"ServerLogin"`
	ServerPlayerId         int    `json:"ServerPlayerId"`
	ConnectionDownloadRate int    `json:"ConnectionDownloadRate"`
	ConnectionUploadRate   int    `json:"ConnectionUploadRate"`
	IsServer               bool   `json:"IsServer"`
	IsDedicated            bool   `json:"IsDedicated"`
}

type TMStatus struct {
	Code int    `json:"Code"`
	Name string `json:"Name"`
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

// Define the wanted api.
func (client *GbxClient) SetApiVersion(version string) error {
	_, err := client.Call("SetApiVersion", version)
	return err
}

// Quit the application. Only available to SuperAdmin.
func (client *GbxClient) QuitGame() error {
	_, err := client.Call("QuitGame")
	return err
}

// Returns the current status of the server.
func (client *GbxClient) GetStatus() (TMStatus, error) {
	res, err := client.Call("GetStatus")
	if err != nil {
		return TMStatus{}, err
	}

	var status TMStatus
	err = convertToStruct(res, &status)
	if err != nil {
		return TMStatus{}, err
	}

	return status, nil
}

type TMVersion struct {
	Name       string `json:"Name"`
	TitleId    string `json:"TitleId"`
	Version    string `json:"Version"`
	Build      string `json:"Build"`
	ApiVersion string `json:"ApiVersion"`
}

// Returns a struct with the Name, TitleId, Version, Build and ApiVersion of the application remotely controlled.
func (client *GbxClient) GetVersion() (TMVersion, error) {
	res, err := client.Call("GetVersion")
	if err != nil {
		return TMVersion{}, err
	}

	var version TMVersion
	err = convertToStruct(res, &version)
	if err != nil {
		return TMVersion{}, err
	}

	return version, nil
}

// Allow the GameServer to call you back.
func (client *GbxClient) EnableCallbacks(enable bool) error {
	_, err := client.Call("EnableCallbacks", enable)
	return err
}
