package main

import "errors"

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
