package main

// Load the black list file with the specified file name. Only available to Admin.
func (client *GbxClient) LoadBlackList(fileName string) error {
	_, err := client.Call("LoadBlackList", fileName)
	return err
}