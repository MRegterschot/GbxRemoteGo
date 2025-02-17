package main

// Allow the GameServer to call you back.
func (client *GbxClient) EnableCallbacks(enable bool) error {
	_, err := client.Send("EnableCallbacks", enable)
	return err
}