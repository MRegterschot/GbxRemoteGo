package main

// Call a vote for a cmd. The command is a XML string corresponding to an XmlRpc request. Only available to Admin.
func (client *GbxClient) CallVote(vote string) error {
	_, err := client.Call("CallVote", vote)
	return err
}