package gbxclient

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

// Returns a boolean if the map with the specified filename matches the current server settings.
func (client *GbxClient) CheckMapForCurrentServerParams(filename string) (bool, error) {
	res, err := client.Call("CheckMapForCurrentServerParams", filename)
	if err != nil {
		return false, err
	}

	data, ok := res.(bool)
	if !ok {
		return false, errors.New("unexpected response format")
	}

	return data, nil
}

// Returns a list of maps among the current selection of the server. This method take two parameters. The first parameter specifies the maximum number of infos to be returned, and the second one the starting index in the selection. The list is an array of structures. Each structure contains the following fields : Name, UId, FileName, Environnement, Author, AuthorNickname, GoldTime, CopperPrice, MapType, MapStyle.
func (client *GbxClient) GetMapList(max int, startIndex int) ([]structs.TMMapInfo, error) {
	res, err := client.Call("GetMapList", max, startIndex)
	if err != nil {
		return nil, err
	}

	var mapList []structs.TMMapInfo
	err = convertToStruct(res, &mapList)
	if err != nil {
		return nil, err
	}

	return mapList, nil
}

// Add the map with the specified filename at the end of the current selection. Only available to Admin.
func (client *GbxClient) AddMap(filename string) error {
	_, err := client.Call("AddMap", filename)
	return err
}

// Add the list of maps with the specified filenames at the end of the current selection. The list of maps to add is an array of strings. Only available to Admin.
func (client *GbxClient) AddMapList(filenames []string) (int, error) {
	res, err := client.Call("AddMapList", filenames)
	if err != nil {
		return 0, err
	}

	data, ok := res.(int)
	if !ok {
		return 0, errors.New("unexpected response format")
	}

	return data, nil
}

// Remove the map with the specified filename from the current selection. Only available to Admin.
func (client *GbxClient) RemoveMap(filename string) error {
	_, err := client.Call("RemoveMap", filename)
	return err
}

// Remove the list of maps with the specified filenames from the current selection. The list of maps to remove is an array of strings. Only available to Admin.
func (client *GbxClient) RemoveMapList(filenames []string) (int, error) {
	res, err := client.Call("RemoveMapList", filenames)
	if err != nil {
		return 0, err
	}

	data, ok := res.(int)
	if !ok {
		return 0, errors.New("unexpected response format")
	}

	return data, nil
}

// Insert the map with the specified filename after the current map. Only available to Admin.
func (client *GbxClient) InsertMap(filename string) error {
	_, err := client.Call("InsertMap", filename)
	return err
}

// Insert the list of maps with the specified filenames after the current map. The list of maps to insert is an array of strings. Only available to Admin.
func (client *GbxClient) InsertMapList(filenames []string) (int, error) {
	res, err := client.Call("InsertMapList", filenames)
	if err != nil {
		return 0, err
	}

	data, ok := res.(int)
	if !ok {
		return 0, errors.New("unexpected response format")
	}

	return data, nil
}

// Set as next map the one with the specified filename, if it is present in the selection. Only available to Admin.
func (client *GbxClient) ChooseNextMap(filename string) error {
	_, err := client.Call("ChooseNextMap", filename)
	return err
}

// Set as next maps the list of maps with the specified filenames, if they are present in the selection. The list of maps to choose is an array of strings. Only available to Admin.
func (client *GbxClient) ChooseNextMapList(filenames []string) (int, error) {
	res, err := client.Call("ChooseNextMapList", filenames)
	if err != nil {
		return 0, err
	}

	data, ok := res.(int)
	if !ok {
		return 0, errors.New("unexpected response format")
	}

	return data, nil
}
