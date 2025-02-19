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
