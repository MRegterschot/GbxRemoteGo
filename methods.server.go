package main

import (
	"errors"

	"github.com/MRegterschot/GbxRemoteGo/structs"
)

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

// Set the download and upload rates (in kbps).
func (client *GbxClient) SetConnectionRates(download int, upload int) error {
	_, err := client.Call("SetConnectionRates", download, upload)
	return err
}

// Returns the list of tags and associated values set on this server. Only available to Admin.
func (client *GbxClient) GetServerTags() ([]structs.TMServerTag, error) {
	res, err := client.Call("GetServerTags")
	if err != nil {
		return nil, err
	}

	var serverTags []structs.TMServerTag
	err = convertToStruct(res, &serverTags)
	if err != nil {
		return nil, err
	}

	return serverTags, nil
}

// Set a tag and its value on the server. This method takes two parameters. The first parameter specifies the name of the tag, and the second one its value. The list is an array of structures {string Name, string Value}. Only available to Admin.
func (client *GbxClient) SetServerTag(tag string, value string) error {
	_, err := client.Call("SetServerTag", tag, value)
	return err
}

// Unset the tag with the specified name on the server. Only available to Admin.
func (client *GbxClient) UnsetServerTag(tag string) error {
	_, err := client.Call("UnsetServerTag", tag)
	return err
}

// Reset all tags on the server. Only available to Admin.
func (client *GbxClient) ResetServerTags() error {
	_, err := client.Call("ResetServerTags")
	return err
}

// Set a new server name in utf8 format. Only available to Admin.
func (client *GbxClient) SetServerName(name string) error {
	_, err := client.Call("SetServerName", name)
	return err
}

// Get the server name in utf8 format.
func (client *GbxClient) GetServerName() (string, error) {
	res, err := client.Call("GetServerName")
	if err != nil {
		return "", err
	}

	// Ensure the response is a string
	data, ok := res.(string)
	if !ok {
		return "", errors.New("unexpected response format")
	}

	return data, nil
}

// Set a new server comment in utf8 format. Only available to Admin.
func (client *GbxClient) SetServerComment(comment string) error {
	_, err := client.Call("SetServerComment", comment)
	return err
}

// Get the server comment in utf8 format.
func (client *GbxClient) GetServerComment() (string, error) {
	res, err := client.Call("GetServerComment")
	if err != nil {
		return "", err
	}

	// Ensure the response is a string
	data, ok := res.(string)
	if !ok {
		return "", errors.New("unexpected response format")
	}

	return data, nil
}

// Set whether the server should be hidden from the public server list (0 = visible, 1 = always hidden, 2 = hidden from nations). Only available to Admin.
func (client *GbxClient) SetHideServer(visibility structs.TMServerVisibility) error {
	_, err := client.Call("SetHideServer", visibility)
	return err
}

// Get whether the server wants to be hidden from the public server list.
func (client *GbxClient) GetHideServer() (structs.TMServerVisibility, error) {
	res, err := client.Call("GetHideServer")
	if err != nil {
		return structs.TMServerVisibility(0), err
	}

	// Ensure the response is an int
	data, ok := res.(int)
	if !ok {
		return structs.TMServerVisibility(0), errors.New("unexpected response format")
	}

	return structs.TMServerVisibility(data), nil
}

// Set a new password for the server. Only available to Admin.
func (client *GbxClient) SetServerPassword(password string) error {
	_, err := client.Call("SetServerPassword", password)
	return err
}

// Get the server password if called as Admin or Super Admin, else returns if a password is needed or not.
func (client *GbxClient) GetServerPassword() (string, error) {
	res, err := client.Call("GetServerPassword")
	if err != nil {
		return "", err
	}

	// Ensure the response is a string
	data, ok := res.(string)
	if !ok {
		return "", errors.New("unexpected response format")
	}

	return data, nil
}

// Set a new password for the spectator mode. Only available to Admin.
func (client *GbxClient) SetServerPasswordForSpectator(password string) error {
	_, err := client.Call("SetServerPasswordForSpectator", password)
	return err
}

// Get the password for spectator mode if called as Admin or Super Admin, else returns if a password is needed or not.
func (client *GbxClient) GetServerPasswordForSpectator() (string, error) {
	res, err := client.Call("GetServerPasswordForSpectator")
	if err != nil {
		return "", err
	}

	// Ensure the response is a string
	data, ok := res.(string)
	if !ok {
		return "", errors.New("unexpected response format")
	}

	return data, nil
}

// Set a new maximum number of players. Only available to Admin. Requires a map restart to be taken into account.
func (client *GbxClient) SetMaxPlayers(maxPlayers int) error {
	_, err := client.Call("SetMaxPlayers", maxPlayers)
	return err
}

// Get the current and next maximum number of players allowed on server. The struct returned contains two fields CurrentValue and NextValue.
func (client *GbxClient) GetMaxPlayers() (structs.TMMaxPlayers, error) {
	res, err := client.Call("GetMaxPlayers")
	if err != nil {
		return structs.TMMaxPlayers{}, err
	}

	var maxPlayers structs.TMMaxPlayers
	err = convertToStruct(res, &maxPlayers)
	if err != nil {
		return structs.TMMaxPlayers{}, err
	}

	return maxPlayers, nil
}

// Set a new maximum number of Spectators. Only available to Admin. Requires a map restart to be taken into account.
func (client *GbxClient) SetMaxSpectators(maxSpectators int) error {
	_, err := client.Call("SetMaxSpectators", maxSpectators)
	return err
}

// Get the current and next maximum number of spectators allowed on server. The struct returned contains two fields CurrentValue and NextValue.
func (client *GbxClient) GetMaxSpectators() (structs.TMMaxSpectators, error) {
	res, err := client.Call("GetMaxSpectators")
	if err != nil {
		return structs.TMMaxSpectators{}, err
	}

	var maxSpectators structs.TMMaxSpectators
	err = convertToStruct(res, &maxSpectators)
	if err != nil {
		return structs.TMMaxSpectators{}, err
	}

	return maxSpectators, nil
}

// Set whether, when a player is switching to spectator, the server should still consider him a player and keep his player slot, or not. Only available to Admin.
func (client *GbxClient) KeepPlayerSlots(keepPlayerSlots bool) error {
	_, err := client.Call("KeepPlayerSlots", keepPlayerSlots)
	return err
}

// Get whether the server keeps player slots when switching to spectator.
func (client *GbxClient) IsKeepingPlayerSlots() (bool, error) {
	res, err := client.Call("IsKeepingPlayerSlots")
	if err != nil {
		return false, err
	}

	// Ensure the response is a bool
	data, ok := res.(bool)
	if !ok {
		return false, errors.New("unexpected response format")
	}

	return data, nil
}

// Returns the path of the game datas directory. Only available to Admin.
func (client *GbxClient) GameDataDirectory() (string, error) {
	res, err := client.Call("GameDataDirectory")
	if err != nil {
		return "", err
	}

	// Ensure the response is a string
	data, ok := res.(string)
	if !ok {
		return "", errors.New("unexpected response format")
	}

	return data, nil
}

// Returns the path of the maps directory. Only available to Admin.
func (client *GbxClient) GetMapsDirectory() (string, error) {
	res, err := client.Call("GetMapsDirectory")
	if err != nil {
		return "", err
	}

	// Ensure the response is a string
	data, ok := res.(string)
	if !ok {
		return "", errors.New("unexpected response format")
	}

	return data, nil
}

// Returns the path of the skins directory. Only available to Admin.
func (client *GbxClient) GetSkinsDirectory() (string, error) {
	res, err := client.Call("GetSkinsDirectory")
	if err != nil {
		return "", err
	}

	// Ensure the response is a string
	data, ok := res.(string)
	if !ok {
		return "", errors.New("unexpected response format")
	}

	return data, nil
}

// Disable player horns. Only available to Admin.
func (client *GbxClient) DisableHorns(disable bool) error {
	_, err := client.Call("DisableHorns", disable)
	return err
}

// Returns whether the horns are disabled.
func (client *GbxClient) AreHornsDisabled() (bool, error) {
	res, err := client.Call("AreHornsDisabled")
	if err != nil {
		return false, err
	}

	// Ensure the response is a bool
	data, ok := res.(bool)
	if !ok {
		return false, errors.New("unexpected response format")
	}

	return data, nil
}
