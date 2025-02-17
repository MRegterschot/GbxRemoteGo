package main

// Quit the application. Only available to SuperAdmin.
func (client *GbxClient) QuitGame() error {
	_, err := client.Send("QuitGame")
	return err
}