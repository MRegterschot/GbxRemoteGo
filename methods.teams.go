package main

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
	data, ok := res.(map[string]interface{})
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
	data, ok := res.(map[string]interface{})
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
