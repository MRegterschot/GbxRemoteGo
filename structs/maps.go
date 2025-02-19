package structs

type TMMapInfo struct {
	Name           string `json:"Name" xmlrpc:"Name"`
	UId            string `json:"UId" xmlrpc:"UId"`
	FileName       string `json:"FileName" xmlrpc:"FileName"`
	Author         string `json:"Author" xmlrpc:"Author"`
	AuthorNickname string `json:"AuthorNickname" xmlrpc:"AuthorNickname"`
	Environnement  string `json:"Environnement" xmlrpc:"Environnement"`
	Mood           string `json:"Mood" xmlrpc:"Mood"`
	BronzeTime     int    `json:"BronzeTime" xmlrpc:"BronzeTime"`
	SilverTime     int    `json:"SilverTime" xmlrpc:"SilverTime"`
	GoldTime       int    `json:"GoldTime" xmlrpc:"GoldTime"`
	AuthorTime     int    `json:"AuthorTime" xmlrpc:"AuthorTime"`
	CopperPrice    int    `json:"CopperPrice" xmlrpc:"CopperPrice"`
	LapRace        bool   `json:"LapRace" xmlrpc:"LapRace"`
	NbLaps         int    `json:"NbLaps" xmlrpc:"NbLaps"`
	NbCheckpoints  int    `json:"NbCheckpoints" xmlrpc:"NbCheckpoints"`
	MapType        string `json:"MapType" xmlrpc:"MapType"`
	MapStyle       string `json:"MapStyle" xmlrpc:"MapStyle"`
}

type TMSMapInfo struct {
	Uid            string `json:"Uid" xmlrpc:"Uid"`
	Name           string `json:"Name" xmlrpc:"Name"`
	FileName       string `json:"FileName" xmlrpc:"FileName"`
	Author         string `json:"Author" xmlrpc:"Author"`
	AuthorNickname string `json:"AuthorNickname" xmlrpc:"AuthorNickname"`
	Environnement  string `json:"Environnement" xmlrpc:"Environnement"`
	Mood           string `json:"Mood" xmlrpc:"Mood"`
	BronzeTime     int    `json:"BronzeTime" xmlrpc:"BronzeTime"`
	SilverTime     int    `json:"SilverTime" xmlrpc:"SilverTime"`
	GoldTime       int    `json:"GoldTime" xmlrpc:"GoldTime"`
	AuthorTime     int    `json:"AuthorTime" xmlrpc:"AuthorTime"`
	CopperPrice    int    `json:"CopperPrice" xmlrpc:"CopperPrice"`
	LapRace        bool   `json:"LapRace" xmlrpc:"LapRace"`
	NbLaps         int    `json:"NbLaps" xmlrpc:"NbLaps"`
	NbCheckpoints  int    `json:"NbCheckpoints" xmlrpc:"NbCheckpoints"`
	MapType        string `json:"MapType" xmlrpc:"MapType"`
	MapStyle       string `json:"MapStyle" xmlrpc:"MapStyle"`
}
