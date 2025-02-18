package main

// Restarts the map with an optional boolean parameter DontClearCupScores
func (client *GbxClient) RestartMap(dontClearCupScores ...bool) error {
	var param interface{} = nil
	if len(dontClearCupScores) > 0 {
		param = dontClearCupScores[0]
	}
	_, err := client.Call("RestartMap", param)
	return err
}

// Switch to next map with an optional boolean parameter DontClearCupScores
func (client *GbxClient) NextMap(dontClearCupScores ...bool) error {
	var param interface{} = nil
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
