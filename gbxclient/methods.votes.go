package gbxclient

import "github.com/MRegterschot/GbxRemoteGo/structs"

// Call a vote for a cmd. The command is a XML string corresponding to an XmlRpc request. Only available to Admin.
func (client *GbxClient) CallVote(vote string) error {
	_, err := client.Call("CallVote", vote)
	return err
}

// Extended call vote. Same as CallVote, but you can additionally supply specific parameters for this vote: a ratio, a time out and who is voting. Special timeout values: a ratio of '-1' means default; a timeout of '0' means default, '1' means indefinite; Voters values: '0' means only active players, '1' means any player, '2' is for everybody, pure spectators included. Only available to Admin.
func (client *GbxClient) CallVoteEx(vote string, ratio float32, timeout int, who int) error {
	_, err := client.Call("CallVoteEx", vote, ratio, timeout, who)
	return err
}

// Cancel the current vote. Only available to Admin.
func (client *GbxClient) CancelVote() error {
	_, err := client.Call("CancelVote")
	return err
}

// Returns the vote currently in progress. The returned structure is { CallerLogin, CmdName, CmdParam }.
func (client *GbxClient) GetCurrentCallVote() (structs.TMCurrentCallVote, error) {
	res, err := client.Call("GetCurrentCallVote")
	if err != nil {
		return structs.TMCurrentCallVote{}, err
	}

	var currentCallVote structs.TMCurrentCallVote
	err = convertToStruct(res, &currentCallVote)
	if err != nil {
		return structs.TMCurrentCallVote{}, err
	}

	return currentCallVote, nil
}
