package gbxclient

import (
	"errors"

	"github.com/MRegterschot/GbxRemoteGo/structs"
)

// Get the current mode script.
func (client *GbxClient) GetModeScriptText() (string, error) {
	res, err := client.Call("GetModeScriptText")
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

// Set the mode script and restart. Only available to Admin.
func (client *GbxClient) SetModeScriptText(script string) error {
	_, err := client.Call("SetModeScriptText", script)
	return err
}

// Returns the description of the current mode script, as a structure containing: Name, CompatibleTypes, Description, Version and the settings available.
func (client *GbxClient) GetModeScriptInfo() (structs.TMModeScriptInfo, error) {
	res, err := client.Call("GetModeScriptInfo")
	if err != nil {
		return structs.TMModeScriptInfo{}, err
	}

	// Ensure the response is a structure
	var modeScriptInfo structs.TMModeScriptInfo
	err = convertToStruct(res, &modeScriptInfo)
	if err != nil {
		return structs.TMModeScriptInfo{}, err
	}

	return modeScriptInfo, nil
}

// Returns the current settings of the mode script.
func (client *GbxClient) GetModeScriptSettings() (map[string]any, error) {
	res, err := client.Call("GetModeScriptSettings")
	if err != nil {
		return nil, err
	}

	// Ensure the response is a map
	data, ok := res.(map[string]any)
	if !ok {
		return nil, errors.New("unexpected response format")
	}

	return data, nil
}

// Change the settings of the mode script. Only available to Admin.
func (client *GbxClient) SetModeScriptSettings(settings map[string]any) error {
	_, err := client.Call("SetModeScriptSettings", settings)
	return err
}

// Send an event to the mode script. Only available to Admin.
func (client *GbxClient) TriggerModeScriptEvent(method string, param string) error {
	_, err := client.Call("TriggerModeScriptEvent", method, param)
	return err
}

// Send an event to the mode script. Only available to Admin.
func (client *GbxClient) TriggerModeScriptEventArray(method string, params []string) error {
	_, err := client.Call("TriggerModeScriptEventArray", method, params)
	return err
}
