package main

import "github.com/MRegterschot/GbxRemoteGo/structs"

// Get some system infos, including connection rates (in kbps).
func (client *GbxClient) GetSystemInfo() (structs.TMSystemInfo, error) {
	res, err := client.Call("GetSystemInfo")
	if err != nil {
		return structs.TMSystemInfo{}, err
	}

	var systemInfo structs.TMSystemInfo
	err = convertToStruct(res, &systemInfo)
	if err != nil {
		return structs.TMSystemInfo{}, err
	}

	return systemInfo, nil
}

// Quit the application. Only available to SuperAdmin.
func (client *GbxClient) QuitGame() error {
	_, err := client.Call("QuitGame")
	return err
}

// Returns the current status of the server.
func (client *GbxClient) GetStatus() (structs.TMStatus, error) {
	res, err := client.Call("GetStatus")
	if err != nil {
		return structs.TMStatus{}, err
	}

	var status structs.TMStatus
	err = convertToStruct(res, &status)
	if err != nil {
		return structs.TMStatus{}, err
	}

	return status, nil
}

// Returns a struct with the Name, TitleId, Version, Build and ApiVersion of the application remotely controlled.
func (client *GbxClient) GetVersion() (structs.TMVersion, error) {
	res, err := client.Call("GetVersion")
	if err != nil {
		return structs.TMVersion{}, err
	}

	var version structs.TMVersion
	err = convertToStruct(res, &version)
	if err != nil {
		return structs.TMVersion{}, err
	}

	return version, nil
}

// Just log the parameters and invoke a callback. Can be used to talk to other xmlrpc clients connected, or to make custom votes. If used in a callvote, the first parameter will be used as the vote message on the clients. Only available to Admin.
func (client *GbxClient) Echo(par1 string, par2 string) error {
	_, err := client.Call("Echo", par1, par2)
	return err
}