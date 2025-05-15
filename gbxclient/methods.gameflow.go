package gbxclient

import "github.com/MRegterschot/GbxRemoteGo/structs"

// Restarts the map with an optional boolean parameter DontClearCupScores
func (client *GbxClient) RestartMap(dontClearCupScores ...bool) error {
	var param any = nil
	if len(dontClearCupScores) > 0 {
		param = dontClearCupScores[0]
	}
	_, err := client.Call("RestartMap", param)
	return err
}

// Switch to next map with an optional boolean parameter DontClearCupScores
func (client *GbxClient) NextMap(dontClearCupScores ...bool) error {
	var param any = nil
	if len(dontClearCupScores) > 0 {
		param = dontClearCupScores[0]
	}
	_, err := client.Call("NextMap", param)
	return err
}

// Attempt to balance teams. Only available to Admin.
func (client *GbxClient) AutoTeamBalance() error {
	_, err := client.Call("AutoTeamBalance")
	return err
}

// Set whether to override the players preferences and always display all opponents (0=no override, 1=show all, other value=minimum number of opponents). Only available to Admin. Requires a map restart to be taken into account.
func (client *GbxClient) SetForceShowAllOpponents(opponents int) error {
	_, err := client.Call("SetForceShowAllOpponents", opponents)
	return err
}

// Get whether players are forced to show all opponents. The struct returned contains two fields CurrentValue and NextValue.
func (client *GbxClient) GetForceShowAllOpponents() (structs.TMForceShowAllOpponents, error) {
	res, err := client.Call("GetForceShowAllOpponents")
	if err != nil {
		return structs.TMForceShowAllOpponents{}, err
	}

	var opponents structs.TMForceShowAllOpponents
	err = convertToStruct(res, &opponents)

	if err != nil {
		return structs.TMForceShowAllOpponents{}, err
	}

	return opponents, nil
}

// Set a new mode script name for script mode. Only available to Admin. Requires a map restart to be taken into account.
func (client *GbxClient) SetScriptName(scriptName string) error {
	_, err := client.Call("SetScriptName", structs.GetScriptByName(scriptName))
	return err
}

// Get the current and next mode script name for script mode. The struct returned contains two fields CurrentValue and NextValue.
func (client *GbxClient) GetScriptName() (structs.TMScriptName, error) {
	res, err := client.Call("GetScriptName")
	if err != nil {
		return structs.TMScriptName{}, err
	}

	var scriptName structs.TMScriptName
	err = convertToStruct(res, &scriptName)

	if err != nil {
		return structs.TMScriptName{}, err
	}

	return scriptName, nil
}
