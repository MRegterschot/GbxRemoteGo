package main

// Load the guest list file with the specified file name. Only available to Admin.
func (client *GbxClient) LoadGuestList(fileName string) error {
	_, err := client.Call("LoadGuestList", fileName)
	return err
}