package main

// Save the guest list in the file with specified file name. Only available to Admin.
func (client *GbxClient) SaveGuestList(fileName string) error {
	_, err := client.Call("SaveGuestList", fileName)
	return err
}