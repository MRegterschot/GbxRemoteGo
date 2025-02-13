package models

type TMSystemInfo struct {
	PublishedIp             string `json:"PublishedIp"`
	Port                    int    `json:"Port"`
	P2PPort                 int    `json:"P2PPort"`
	TitleId                 string `json:"TitleId"`
	ServerLogin             string `json:"ServerLogin"`
	ServerPlayerId          int    `json:"ServerPlayerId"`
	ConnectionDownloadRate  int    `json:"ConnectionDownloadRate"`
	ConnectionUploadRate    int    `json:"ConnectionUploadRate"`
	IsServer                bool   `json:"IsServer"`
	IsDedicated             bool   `json:"IsDedicated"`
}