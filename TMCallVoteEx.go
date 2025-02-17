package main

// Extended call vote. Same as CallVote, but you can additionally supply specific parameters for this vote: a ratio, a time out and who is voting. Special timeout values: a ratio of '-1' means default; a timeout of '0' means default, '1' means indefinite; Voters values: '0' means only active players, '1' means any player, '2' is for everybody, pure spectators included. Only available to Admin.
func (client *GbxClient) CallVoteEx(vote string, ratio float32, timeout int, who int) error {
	_, err := client.Send("CallVoteEx", vote, ratio, timeout, who)
	return err
}