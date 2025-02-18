package structs

type BanList struct {
	Login      string `json:"Login" xmlrpc:"Login"`
	ClientName string `json:"ClientName" xmlrpc:"ClientName"`
	IPAddress  string `json:"IPAddress" xmlrpc:"IPAddress"`
}
