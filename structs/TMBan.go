package structs

type BanList struct {
	Login      string `json:"Login"`
	ClientName string `json:"ClientName"`
	IPAddress  string `json:"IPAddress"`
}
