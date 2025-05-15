package events

import "github.com/MRegterschot/GbxRemoteGo/structs"

type MapEventArgs struct {
	Map structs.TMSMapInfo `json:"Map" xmlrpc:"Map"`
}

type EndMatchEventArgs struct {
	Rankings   []structs.TMSPlayerRanking `json:"Rankings" xmlrpc:"Rankings"`
	WinnerTeam int                        `json:"WinnerTeam" xmlrpc:"WinnerTeam"`
}

type MapListModifiedEventArgs struct {
	CurMapIndex    int  `json:"CurMapIndex" xmlrpc:"CurMapIndex"`
	NextMapIndex   int  `json:"NextMapIndex" xmlrpc:"NextMapIndex"`
	IsListModified bool `json:"IsListModified" xmlrpc:"IsListModified"`
}

type VoteUpdatedEventArgs struct {
	Login     string `json:"Login" xmlrpc:"Login"`
	StateName string `json:"StateName" xmlrpc:"StateName"`
	CmdName   string `json:"CmdName" xmlrpc:"CmdName"`
	CmdParam  string `json:"CmdParam" xmlrpc:"CmdParam"`
}

type ScoresEventArgs struct {
	ResponseId   string       `json:"responseid"`
	Section      string       `json:"section"`
	UseTeams     bool         `json:"useteams"`
	WinnerTeam   int          `json:"winnerteam"`
	WinnerPlayer string       `json:"winnerplayer"`
	Teams        []TeamArgs   `json:"teams"`
	Players      []PlayerArgs `json:"players"`
}

type TeamArgs struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	RoundPoints int    `json:"roundpoints"`
	MapPoints   int    `json:"mappoints"`
	MatchPoints int    `json:"matchpoints"`
}

type PlayerArgs struct {
	Login               string `json:"login"`
	AccountId           string `json:"accountid"`
	Name                string `json:"name"`
	Team                int    `json:"team"`
	Rank                int    `json:"rank"`
	RoundPoints         int    `json:"roundpoints"`
	MapPoints           int    `json:"mappoints"`
	MatchPoints         int    `json:"matchpoints"`
	BestRaceTime        int    `json:"bestracetime"`
	BestRaceCheckpoints []int  `json:"bestracecheckpoints"`
	BestLapTime         int    `json:"bestlaptime"`
	BestLapCheckpoints  []int  `json:"bestlapcheckpoints"`
	PrevRaceTime        int    `json:"prevracetime"`
	PrevRaceCheckpoints []int  `json:"prevracecheckpoints"`
}

type StartLineEventArgs struct {
	Time      int    `json:"time" xmlrpc:"Time"`
	AccountId string `json:"accountid" xmlrpc:"AccountId"`
	Login     string `json:"login" xmlrpc:"Login"`
}

type WarmUpEventArgs struct {
	Current int `json:"current" xmlrpc:"Current"`
	Total   int `json:"total" xmlrpc:"Total"`
}
