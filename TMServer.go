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