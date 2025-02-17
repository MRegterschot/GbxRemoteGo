package main

type TMVersion struct {
	Name 	string `json:"Name"`
	TitleId string `json:"TitleId"`
	Version string `json:"Version"`
	Build  string `json:"Build"`
	ApiVersion string `json:"ApiVersion"`
}

// Returns a struct with the Name, TitleId, Version, Build and ApiVersion of the application remotely controlled.
func (client *GbxClient) GetVersion() (TMVersion, error) {
	res, err := client.Call("GetVersion")
	if err != nil {
		return TMVersion{}, err
	}

	var version TMVersion
	err = convertToStruct(res, &version)
	if err != nil {
		return TMVersion{}, err
	}

	return version, nil
}