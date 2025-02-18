package main

// Save the black list in the file with specified file name. Only available to Admin.
func (client *GbxClient) SaveBlackList(fileName string) error {
	_, err := client.Call("SaveBlackList", fileName)
	return err
}