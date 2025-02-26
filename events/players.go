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

type PlayerWayPointEventArgs struct {
	Time               int     `json:"time" xmlrpc:"Time"`
	Login              string  `json:"login" xmlrpc:"Login"`
	AccountId          string  `json:"accountid" xmlrpc:"AccountId"`
	RaceTime           int     `json:"racetime" xmlrpc:"RaceTime"`
	LapTime            int     `json:"laptime" xmlrpc:"LapTime"`
	CheckpointInRace   int     `json:"checkpointinrace" xmlrpc:"CheckpointInRace"`
	CheckpointInLap    int     `json:"checkpointinlap" xmlrpc:"CheckpointInLap"`
	IsEndRace          bool    `json:"isendrace" xmlrpc:"IsEndRace"`
	IsEndLap           bool    `json:"isendlap" xmlrpc:"IsEndLap"`
	IsInfiniteLaps     bool    `json:"isinfinitelaps" xmlrpc:"IsInfiniteLaps"`
	IsIndependentLaps  bool    `json:"isindependentlaps" xmlrpc:"IsIndependentLaps"`
	CurRaceCheckpoints []int   `json:"curracecheckpoints" xmlrpc:"CurRaceCheckpoints"`
	CurLapCheckpoints  []int   `json:"curlapcheckpoints" xmlrpc:"CurLapCheckpoints"`
	BlockId            string  `json:"blockid" xmlrpc:"BlockId"`
	Speed              float64 `json:"speed" xmlrpc:"Speed"`
}

type PlayerIncoherenceEventArgs struct {
	Login     string `json:"Login" xmlrpc:"Login"`
	PlayerUid int    `json:"PlayerUid" xmlrpc:"PlayerUid"`
}
