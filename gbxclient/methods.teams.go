package gbxclient

import (
	"errors"

	"github.com/MRegterschot/GbxRemoteGo/structs"
)

// Set Team names and colors (deprecated). Only available to Admin.
func (client *GbxClient) SetTeamInfo(name string, huePrimary float32, zonePath string, par4 string, par5 float32, par6 string, par7 string, par8 float32, par9 string) error {
	_, err := client.Call("SetTeamInfo", name, huePrimary, zonePath, par4, par5, par6, par7, par8, par9)
	return err
}

// Return Team info for a given clan (0 = no clan, 1, 2). The structure contains: Name, ZonePath, City, EmblemUrl, HuePrimary, HueSecondary, RGB, ClubLinkUrl. Only available to Admin.
func (client *GbxClient) GetTeamInfo(clanID int) (structs.TMTeamInfo, error) {
	res, err := client.Call("GetTeamInfo", clanID)
	if err != nil {
		return structs.TMTeamInfo{}, err
	}

	// Ensure the response is a struct
	data, ok := res.(map[string]any)
	if !ok {
		return structs.TMTeamInfo{}, errors.New("unexpected response format")
	}

	// Convert struct to TMTeamInfo
	var team structs.TMTeamInfo
	err = convertToStruct(data, &team)
	if err != nil {
		return structs.TMTeamInfo{}, err
	}

	return team, nil
}

// Set the clublinks to use for the two clans. Only available to Admin.
func (client *GbxClient) SetForcedClubLinks(clubLink1 string, clubLink2 string) error {
	_, err := client.Call("SetForcedClubLinks", clubLink1, clubLink2)
	return err
}

// Get the forced clublinks.
func (client *GbxClient) GetForcedClubLinks() (structs.TMForcedClubLinks, error) {
	res, err := client.Call("GetForcedClubLinks")
	if err != nil {
		return structs.TMForcedClubLinks{}, err
	}

	// Ensure the response is a struct
	data, ok := res.(map[string]any)
	if !ok {
		return structs.TMForcedClubLinks{}, errors.New("unexpected response format")
	}

	// Convert struct to TMForcedClubLinks
	var links structs.TMForcedClubLinks
	err = convertToStruct(data, &links)
	if err != nil {
		return structs.TMForcedClubLinks{}, err
	}

	return links, nil
}

// Set whether the players can choose their side or if the teams are forced by the server (using ForcePlayerTeam()). Only available to Admin.
func (client *GbxClient) SetForcedTeams(forced bool) error {
	_, err := client.Call("SetForcedTeams", forced)
	return err
}

// Returns whether the players can choose their side or if the teams are forced by the server.
func (client *GbxClient) GetForcedTeams() (bool, error) {
	res, err := client.Call("GetForcedTeams")
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

// Returns the current winning team for the race in progress. (-1: if not in team mode, or draw match)
func (client *GbxClient) GetCurrentWinnerTeam() (int, error) {
	res, err := client.Call("GetCurrentWinnerTeam")
	if err != nil {
		return 0, err
	}

	data, ok := res.(int)
	if !ok {
		return 0, errors.New("unexpected response format")
	}

	return data, nil
}

// Force the team of the player. Only available in team mode. You have to pass the login and the team number (0 or 1). Only available to Admin.
func (client *GbxClient) ForcePlayerTeam(login string, team int) error {
	_, err := client.Call("ForcePlayerTeam", login, team)
	return err
}

// Force the team of the player. Only available in team mode. You have to pass the playerid and the team number (0 or 1). Only available to Admin.
func (client *GbxClient) ForcePlayerTeamId(playerID int, team int) error {
	_, err := client.Call("ForcePlayerTeamId", playerID, team)
	return err
}
