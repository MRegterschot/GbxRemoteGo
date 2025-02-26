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
	Time               int     `json:"Time" xmlrpc:"Time"`
	Login              string  `json:"Login" xmlrpc:"Login"`
	AccountId          string  `json:"AccountId" xmlrpc:"AccountId"`
	RaceTime           int     `json:"RaceTime" xmlrpc:"RaceTime"`
	LapTime            int     `json:"LapTime" xmlrpc:"LapTime"`
	CheckpointInRace   int     `json:"CheckpointInRace" xmlrpc:"CheckpointInRace"`
	CheckpointInLap    int     `json:"CheckpointInLap" xmlrpc:"CheckpointInLap"`
	IsEndRace          bool    `json:"IsEndRace" xmlrpc:"IsEndRace"`
	IsEndLap           bool    `json:"IsEndLap" xmlrpc:"IsEndLap"`
	IsInfiniteLaps     bool    `json:"IsInfiniteLaps" xmlrpc:"IsInfiniteLaps"`
	IsIndependentLaps  bool    `json:"IsIndependentLaps" xmlrpc:"IsIndependentLaps"`
	CurRaceCheckpoints []int   `json:"CurRaceCheckpoints" xmlrpc:"CurRaceCheckpoints"`
	CurLapCheckpoints  []int   `json:"CurLapCheckpoints" xmlrpc:"CurLapCheckpoints"`
	BlockId            string  `json:"BlockId" xmlrpc:"BlockId"`
	Speed              float64 `json:"Speed" xmlrpc:"Speed"`
}

type PlayerIncoherenceEventArgs struct {
	Login     string `json:"Login" xmlrpc:"Login"`
	PlayerUid int    `json:"PlayerUid" xmlrpc:"PlayerUid"`
}
