package main

import (
	"errors"

	"github.com/MRegterschot/GbxRemoteGo/structs"
)

// Returns the current map index in the selection, or -1 if the map is no longer in the selection.
func (client *GbxClient) GetCurrentMapIndex() (int, error) {
	res, err := client.Call("GetCurrentMapIndex")
	if err != nil {
		return 0, err
	}

	data, ok := res.(int)
	if !ok {
		return 0, errors.New("unexpected response format")
	}

	return data, nil
}

// Returns the map index in the selection that will be played next (unless the current one is restarted...)
func (client *GbxClient) GetNextMapIndex() (int, error) {
	res, err := client.Call("GetNextMapIndex")
	if err != nil {
		return 0, err
	}

	data, ok := res.(int)
	if !ok {
		return 0, errors.New("unexpected response format")
	}

	return data, nil
}

// Sets the map index in the selection that will be played next (unless the current one is restarted...)
func (client *GbxClient) SetNextMapIndex(index int) error {
	_, err := client.Call("SetNextMapIndex", index)
	return err
}

// Sets the map in the selection that will be played next (unless the current one is restarted...)
func (client *GbxClient) SetNextMapIdent(id string) error {
	_, err := client.Call("SetNextMapIdent", id)
	return err
}

// Immediately jumps to the map designated by the index in the selection.
func (client *GbxClient) JumpToMapIndex(index int) error {
	_, err := client.Call("JumpToMapIndex", index)
	return err
}

// Immediately jumps to the map designated by its identifier (it must be in the selection).
func (client *GbxClient) JumpToMapIdent(id string) error {
	_, err := client.Call("JumpToMapIdent", id)
	return err
}

// Returns a struct containing the infos for the current map. The struct contains the following fields : Name, UId, FileName, Author, AuthorNickname, Environnement, Mood, BronzeTime, SilverTime, GoldTime, AuthorTime, CopperPrice, LapRace, NbLaps, NbCheckpoints, MapType, MapStyle.
func (client *GbxClient) GetCurrentMapInfo() (structs.TMMapInfo, error) {
	res, err := client.Call("GetCurrentMapInfo")
	if err != nil {
		return structs.TMMapInfo{}, err
	}

	var mapInfo structs.TMMapInfo
	err = convertToStruct(res, &mapInfo)
	if err != nil {
		return structs.TMMapInfo{}, err
	}

	return mapInfo, nil
}

// Returns a struct containing the infos for the next map. The struct contains the following fields : Name, UId, FileName, Author, AuthorNickname, Environnement, Mood, BronzeTime, SilverTime, GoldTime, AuthorTime, CopperPrice, LapRace, MapType, MapStyle. (NbLaps and NbCheckpoints are also present but always set to -1)
func (client *GbxClient) GetNextMapInfo() (structs.TMMapInfo, error) {
	res, err := client.Call("GetNextMapInfo")
	if err != nil {
		return structs.TMMapInfo{}, err
	}

	var mapInfo structs.TMMapInfo
	err = convertToStruct(res, &mapInfo)
	if err != nil {
		return structs.TMMapInfo{}, err
	}

	return mapInfo, nil
}

// Returns a struct containing the infos for the map with the specified filename. The struct contains the following fields : Name, UId, FileName, Author, AuthorNickname, Environnement, Mood, BronzeTime, SilverTime, GoldTime, AuthorTime, CopperPrice, LapRace, MapType, MapStyle. (NbLaps and NbCheckpoints are also present but always set to -1)
func (client *GbxClient) GetMapInfo(filename string) (structs.TMMapInfo, error) {
	res, err := client.Call("GetMapInfo", filename)
	if err != nil {
		return structs.TMMapInfo{}, err
	}

	var mapInfo structs.TMMapInfo
	err = convertToStruct(res, &mapInfo)
	if err != nil {
		return structs.TMMapInfo{}, err
	}

	return mapInfo, nil
}
