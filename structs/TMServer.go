package structs

type TMSystemInfo struct {
	PublishedIp            string `json:"PublishedIp"`
	Port                   int    `json:"Port"`
	P2PPort                int    `json:"P2PPort"`
	TitleId                string `json:"TitleId"`
	ServerLogin            string `json:"ServerLogin"`
	ServerPlayerId         int    `json:"ServerPlayerId"`
	ConnectionDownloadRate int    `json:"ConnectionDownloadRate"`
	ConnectionUploadRate   int    `json:"ConnectionUploadRate"`
	IsServer               bool   `json:"IsServer"`
	IsDedicated            bool   `json:"IsDedicated"`
}

type TMStatus struct {
	Code int    `json:"Code"`
	Name string `json:"Name"`
}

type TMVersion struct {
	Name       string `json:"Name"`
	TitleId    string `json:"TitleId"`
	Version    string `json:"Version"`
	Build      string `json:"Build"`
	ApiVersion string `json:"ApiVersion"`
}

type TMServerTag struct {
	Name  string `json:"Name" xmlrpc:"Name"`
	Value string `json:"Value" xmlrpc:"Value"`
}

type TMServerVisibility int

const (
	TMServerVisible TMServerVisibility = iota
	TMServerHidden
	TMServerHiddenFromNations
)

type TMMaxPlayers struct {
	CurrentValue int `json:"CurrentValue" xmlrpc:"CurrentValue"`
	NextValue    int `json:"NextValue" xmlrpc:"NextValue"`
}

type TMMaxSpectators struct {
	CurrentValue int `json:"CurrentValue" xmlrpc:"CurrentValue"`
	NextValue    int `json:"NextValue" xmlrpc:"NextValue"`
}
