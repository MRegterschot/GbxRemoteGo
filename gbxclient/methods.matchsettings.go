package gbxclient

import "errors"

// Set a list of maps defined in the playlist with the specified filename as the current selection of the server, and load the gameinfos from the same file. Only available to Admin.
func (client *GbxClient) LoadMatchSettings(filename string) (int, error) {
	res, err := client.Call("LoadMatchSettings", filename)
	if err != nil {
		return 0, err
	}

	data, ok := res.(int)
	if !ok {
		return 0, errors.New("unexpected response format")
	}

	return data, nil
}

// Add a list of maps defined in the playlist with the specified filename at the end of the current selection. Only available to Admin.
func (client *GbxClient) AppendPlaylistFromMatchSettings(filename string) (int, error) {
	res, err := client.Call("AppendPlaylistFromMatchSettings", filename)
	if err != nil {
		return 0, err
	}

	data, ok := res.(int)
	if !ok {
		return 0, errors.New("unexpected response format")
	}

	return data, nil
}

// Save the current selection of map in the playlist with the specified filename, as well as the current gameinfos. Only available to Admin.
func (client *GbxClient) SaveMatchSettings(filename string) (int, error) {
	res, err := client.Call("SaveMatchSettings", filename)
	if err != nil {
		return 0, err
	}

	data, ok := res.(int)
	if !ok {
		return 0, errors.New("unexpected response format")
	}

	return data, nil
}

// Insert a list of maps defined in the playlist with the specified filename after the current map. Only available to Admin.
func (client *GbxClient) InsertPlaylistFromMatchSettings(filename string) (int, error) {
	res, err := client.Call("InsertPlaylistFromMatchSettings", filename)
	if err != nil {
		return 0, err
	}

	data, ok := res.(int)
	if !ok {
		return 0, errors.New("unexpected response format")
	}

	return data, nil
}