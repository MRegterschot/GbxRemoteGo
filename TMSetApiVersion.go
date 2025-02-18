package main

// Define the wanted api.
func (client *GbxClient) SetApiVersion(version string) error {
	_, err := client.Call("SetApiVersion", version)
	return err
}