package main

type TMCurrentCallVote struct {
	CallerLogin string `json:"CallerLogin"`
	CmdName     string `json:"CmdName"`
	CmdParam    string `json:"CmdParam"`
}

// Returns the vote currently in progress. The returned structure is { CallerLogin, CmdName, CmdParam }.
func (client *GbxClient) GetCurrentCallVote() (TMCurrentCallVote, error) {
	res, err := client.Call("GetCurrentCallVote")
	if err != nil {
		return TMCurrentCallVote{}, err
	}

	var currentCallVote TMCurrentCallVote
	err = convertToStruct(res, &currentCallVote)
	if err != nil {
		return TMCurrentCallVote{}, err
	}

	return currentCallVote, nil
}
