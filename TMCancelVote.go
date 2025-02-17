package main
// Cancel the current vote. Only available to Admin.
func (client *GbxClient) CancelVote() error {
	_, err := client.Send("CancelVote")
	return err
}