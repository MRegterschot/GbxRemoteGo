package structs

type TMTeamInfo struct {
	Name         string  `json:"Name" xmlrpc:"Name"`
	ZonePath     string  `json:"ZonePath" xmlrpc:"ZonePath"`
	City         string  `json:"City" xmlrpc:"City"`
	EmblemUrl    string  `json:"EmblemUrl" xmlrpc:"EmblemUrl"`
	ClubLinkUrl  string  `json:"ClubLinkUrl" xmlrpc:"ClubLinkUrl"`
	HuePrimary   float32 `json:"HuePrimary" xmlrpc:"HuePrimary"`
	HueSecondary float32 `json:"HueSecondary" xmlrpc:"HueSecondary"`
	RGB          string  `json:"RGB" xmlrpc:"RGB"`
}

type TMForcedClubLinks struct {
	ClubLink1 string `json:"ClubLink1" xmlrpc:"ClubLink1"`
	ClubLink2 string `json:"ClubLink2" xmlrpc:"ClubLink2"`
}
