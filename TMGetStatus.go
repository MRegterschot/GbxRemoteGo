package main

type TMStatus struct {
	Code int `json:"Code"`
	Name string `json:"Name"`
}

// Returns the current status of the server.
func (client *GbxClient) GetStatus() (TMStatus, error) {
	res, err := client.Call("GetStatus")
	if err != nil {
		return TMStatus{}, err
	}

	var status TMStatus
	err = convertToStruct(res, &status)
	if err != nil {
		return TMStatus{}, err
	}

	return status, nil
}