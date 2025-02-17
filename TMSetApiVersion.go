package main

// Define the wanted api.
func (client *GbxClient) SetApiVersion(version string) error {
	_, err := client.Send("SetApiVersion", version)
	return err
}