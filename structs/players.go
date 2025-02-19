package structs

type TMBanListEntry struct {
	Login      string `json:"Login" xmlrpc:"Login"`
	ClientName string `json:"ClientName" xmlrpc:"ClientName"`
	IPAddress  string `json:"IPAddress" xmlrpc:"IPAddress"`
}

type TMBlackListEntry struct {
	Login string `json:"Login" xmlrpc:"Login"`
}

type TMGuestListEntry struct {
	Login string `json:"Login" xmlrpc:"Login"`
}

type TMPlayerInfo struct {
	Login           string  `json:"Login" xmlrpc:"Login"`
	NickName        string  `json:"NickName" xmlrpc:"NickName"`
	PlayerId        int     `json:"PlayerId" xmlrpc:"PlayerId"`
	TeamId          int     `json:"TeamId" xmlrpc:"TeamId"`
	SpectatorStatus int     `json:"SpectatorStatus" xmlrpc:"SpectatorStatus"`
	LadderRanking   int     `json:"LadderRanking" xmlrpc:"LadderRanking"`
	Flags           int     `json:"Flags" xmlrpc:"Flags"`
	LadderScore     float32 `json:"LadderScore" xmlrpc:"LadderScore"`
}

type TMPlayerDetailedInfo struct {
	Login                     string        `json:"Login" xmlrpc:"Login"`
	NickName                  string        `json:"NickName" xmlrpc:"NickName"`
	PlayerId                  int           `json:"PlayerId" xmlrpc:"PlayerId"`
	TeamId                    int           `json:"TeamId" xmlrpc:"TeamId"`
	Path                      string        `json:"Path" xmlrpc:"Path"`
	Language                  string        `json:"Language" xmlrpc:"Language"`
	ClientVersion             string        `json:"ClientVersion" xmlrpc:"ClientVersion"`
	ClientTitleVersion        string        `json:"ClientTitleVersion" xmlrpc:"ClientTitleVersion"`
	IPAddress                 string        `json:"IPAddress" xmlrpc:"IPAddress"`
	DownloadRate              int           `json:"DownloadRate" xmlrpc:"DownloadRate"`
	UploadRate                int           `json:"UploadRate" xmlrpc:"UploadRate"`
	IsSpectator               bool          `json:"IsSpectator" xmlrpc:"IsSpectator"`
	IsInOfficialMode          bool          `json:"IsInOfficialMode" xmlrpc:"IsInOfficialMode"`
	IsReferee                 bool          `json:"IsReferee" xmlrpc:"IsReferee"`
	Avatar                    TMAvatar      `json:"Avatar" xmlrpc:"Avatar"`
	Skins                     []TMSkin      `json:"Skins" xmlrpc:"Skins"`
	LadderStats               TMLadderStats `json:"LadderStats" xmlrpc:"LadderStats"`
	HoursSinceZoneInscription int           `json:"HoursSinceZoneInscription" xmlrpc:"HoursSinceZoneInscription"`
	BroadcasterLogin          string        `json:"BroadcasterLogin" xmlrpc:"BroadcasterLogin"`
	Allies                    []string      `json:"Allies" xmlrpc:"Allies"`
	ClubLink                  string        `json:"ClubLink" xmlrpc:"ClubLink"`
}

type TMAvatar struct {
	FileName string `json:"FileName" xmlrpc:"FileName"`
	Checksum string `json:"Checksum" xmlrpc:"Checksum"`
}

type TMSkin struct {
	Environnement string     `json:"Environnement" xmlrpc:"Environnement"`
	PackDesc      TMPackDesc `json:"PackDesc" xmlrpc:"PackDesc"`
}

type TMPackDesc struct {
	FileName string `json:"FileName" xmlrpc:"FileName"`
	Checksum string `json:"Checksum" xmlrpc:"Checksum"`
}

type TMLadderStats struct {
	LastMatchScore float32         `json:"LastMatchScore" xmlrpc:"LastMatchScore"`
	NbrMatchWins   int             `json:"NbrMatchWins" xmlrpc:"NbrMatchWins"`
	NbrMatchDraws  int             `json:"NbrMatchDraws" xmlrpc:"NbrMatchDraws"`
	NbrMatchLosses int             `json:"NbrMatchLosses" xmlrpc:"NbrMatchLosses"`
	TeamName       string          `json:"TeamName" xmlrpc:"TeamName"`
	PlayerRankings []TMZoneRanking `json:"PlayerRankings" xmlrpc:"PlayerRankings"`
	TeamRankings   []interface{}   `json:"TeamRankings" xmlrpc:"TeamRankings"`
}

type TMZoneRanking struct {
	Path       string  `json:"Path" xmlrpc:"Path"`
	Score      float32 `json:"Score" xmlrpc:"Score"`
	Ranking    int     `json:"Ranking" xmlrpc:"Ranking"`
	TotalCount int     `json:"TotalCount" xmlrpc:"TotalCount"`
}

type TMPlayerRanking struct {
	Login    string `json:"Login" xmlrpc:"Login"`
	NickName string `json:"NickName" xmlrpc:"NickName"`
	PlayerId int    `json:"PlayerId" xmlrpc:"PlayerId"`
	Rank     int    `json:"Rank" xmlrpc:"Rank"`

	BestTime        int     `json:"BestTime" xmlrpc:"BestTime"`               // Deprecated
	BestCheckpoints []int   `json:"BestCheckpoints" xmlrpc:"BestCheckpoints"` // Deprecated
	Score           int     `json:"Score" xmlrpc:"Score"`                     // Deprecated
	NbrLapsFinished int     `json:"NbrLapsFinished" xmlrpc:"NbrLapsFinished"` // Deprecated
	LadderScore     float32 `json:"LadderScore" xmlrpc:"LadderScore"`         // Deprecated
}

type TMPlayerScore struct {
	PlayerId int `json:"PlayerId" xmlrpc:"PlayerId"`
	Score    int `json:"Score" xmlrpc:"Score"`
}

type TMSPlayerInfo struct {
	Login           string `json:"Login" xmlrpc:"Login"`
	NickName        string `json:"NickName" xmlrpc:"NickName"`
	PlayerId        int    `json:"PlayerId" xmlrpc:"PlayerId"`
	TeamId          int    `json:"TeamId" xmlrpc:"TeamId"`
	SpectatorStatus int    `json:"SpectatorStatus" xmlrpc:"SpectatorStatus"`
	LadderRanking   int    `json:"LadderRanking" xmlrpc:"LadderRanking"`
	Flags           int    `json:"Flags" xmlrpc:"Flags"`
}

type TMSPlayerRanking struct {
	Login           string  `json:"Login" xmlrpc:"Login"`
	NickName        string  `json:"NickName" xmlrpc:"NickName"`
	PlayerId        int     `json:"PlayerId" xmlrpc:"PlayerId"`
	Rank            int     `json:"Rank" xmlrpc:"Rank"`
	BestTime        int     `json:"BestTime" xmlrpc:"BestTime"`
	BestCheckpoints []int   `json:"BestCheckpoints" xmlrpc:"BestCheckpoints"`
	Score           int     `json:"Score" xmlrpc:"Score"`
	NbrLapsFinished int     `json:"NbrLapsFinished" xmlrpc:"NbrLapsFinished"`
	LadderScore     float32 `json:"LadderScore" xmlrpc:"LadderScore"`
}
