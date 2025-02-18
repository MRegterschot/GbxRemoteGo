package main
// Cancel the current vote. Only available to Admin.
func (client *GbxClient) CancelVote() error {
	_, err := client.Call("CancelVote")
	return err
}