package events

import "github.com/MRegterschot/GbxRemoteGo/structs"

type PlayerConnectEventArgs struct {
	Login       string `json:"Login" xmlrpc:"Login"`
	IsSpectator bool   `json:"IsSpectator" xmlrpc:"IsSpectator"`
}

type PlayerDisconnectEventArgs struct {
	Login  string `json:"Login" xmlrpc:"Login"`
	Reason string `json:"Reason" xmlrpc:"Reason"`
}

type PlayerChatEventArgs struct {
	Login          string `json:"Login" xmlrpc:"Login"`
	Text           string `json:"Text" xmlrpc:"Text"`
	PlayerUid      int    `json:"PlayerUid" xmlrpc:"PlayerUid"`
	IsRegistredCmd bool   `json:"IsRegistredCmd" xmlrpc:"IsRegistredCmd"`
	Options        int    `json:"Options" xmlrpc:"Options"`
}

type PlayerAlliesChangedEventArgs struct {
	Login string `json:"Login" xmlrpc:"Login"`
}

type PlayerInfoChangedEventArgs struct {
	PlayerInfo structs.TMSPlayerInfo `json:"PlayerInfo" xmlrpc:"PlayerInfo"`
}

type PlayerManialinkPageAnswerEventArgs struct {
	Login     string                `json:"Login" xmlrpc:"Login"`
	Answer    string                `json:"Answer" xmlrpc:"Answer"`
	PlayerUid int                   `json:"PlayerUid" xmlrpc:"PlayerUid"`
	Entries   []structs.TMSEntryVal `json:"Entries" xmlrpc:"Entries"`
}

type PlayerCheckpointEventArgs struct {
	Login           string `json:"Login" xmlrpc:"Login"`
	PlayerUid       int    `json:"PlayerUid" xmlrpc:"PlayerUid"`
	TimeOrScore     int    `json:"TimeOrScore" xmlrpc:"TimeOrScore"`
	CurLap          int    `json:"CurLap" xmlrpc:"CurLap"`
	CheckpointIndex int    `json:"CheckpointIndex" xmlrpc:"CheckpointIndex"`
}

type PlayerFinishEventArgs struct {
	Login       string `json:"Login" xmlrpc:"Login"`
	PlayerUid   int    `json:"PlayerUid" xmlrpc:"PlayerUid"`
	TimeOrScore int    `json:"TimeOrScore" xmlrpc:"TimeOrScore"`
}

type PlayerIncoherenceEventArgs struct {
	Login     string `json:"Login" xmlrpc:"Login"`
	PlayerUid int    `json:"PlayerUid" xmlrpc:"PlayerUid"`
}
