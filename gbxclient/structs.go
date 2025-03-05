package gbxclient

import (
	"bytes"
	"net"
	"sync"

	"github.com/MRegterschot/GbxRemoteGo/events"
)

type Options struct {
	ShowErrors  bool
	ThrowErrors bool
}

type GbxCallbackStruct[T any] struct {
	Key  string
	Call func(*GbxClient, T)
}

type GbxCallbacks struct {
	OnAnyCallback               []GbxCallbackStruct[CallbackEventArgs]
	OnBeginMap                  []GbxCallbackStruct[events.MapEventArgs]
	OnBeginMatch                []GbxCallbackStruct[struct{}] // No args
	OnEcho                      []GbxCallbackStruct[events.EchoEventArgs]
	OnEndMap                    []GbxCallbackStruct[events.MapEventArgs]
	OnEndMatch                  []GbxCallbackStruct[events.EndMatchEventArgs]
	OnEndRound                  []GbxCallbackStruct[events.ScoresEventArgs]
	OnPreEndRound               []GbxCallbackStruct[events.ScoresEventArgs]
	OnMapListModified           []GbxCallbackStruct[events.MapListModifiedEventArgs]
	OnPlayerAlliesChanged       []GbxCallbackStruct[events.PlayerAlliesChangedEventArgs]
	OnPlayerChat                []GbxCallbackStruct[events.PlayerChatEventArgs]
	OnPlayerConnect             []GbxCallbackStruct[events.PlayerConnectEventArgs]
	OnPlayerDisconnect          []GbxCallbackStruct[events.PlayerDisconnectEventArgs]
	OnPlayerInfoChanged         []GbxCallbackStruct[events.PlayerInfoChangedEventArgs]
	OnPlayerManialinkPageAnswer []GbxCallbackStruct[events.PlayerManialinkPageAnswerEventArgs]
	OnServerStart               []GbxCallbackStruct[struct{}] // No args
	OnServerStop                []GbxCallbackStruct[struct{}] // No args
	OnStatusChanged             []GbxCallbackStruct[events.StatusChangedEventArgs]
	OnTunnelDataReceived        []GbxCallbackStruct[events.TunnelDataReceivedEventArgs]
	OnVoteUpdated               []GbxCallbackStruct[events.VoteUpdatedEventArgs]
	OnPlayerCheckpoint          []GbxCallbackStruct[events.PlayerWayPointEventArgs]
	OnPlayerFinish              []GbxCallbackStruct[events.PlayerWayPointEventArgs]
	OnPlayerIncoherence         []GbxCallbackStruct[events.PlayerIncoherenceEventArgs]
}

type GbxClient struct {
	GbxCallbacks
	Host             string
	Port             int
	IsConnected      bool
	DoHandShake      bool
	ReqHandle        uint32
	Socket           net.Conn
	Mutex            sync.Mutex
	RecvData         bytes.Buffer
	ResponseLength   *uint32
	DataPointer      int
	Options          Options
	PromiseCallbacks map[uint32]chan PromiseResult
	Events           EventEmitter
}

type PromiseResult struct {
	Result interface{}
	Error  error
}

type EventEmitter struct {
	events map[string][]chan interface{}
	mu     sync.Mutex
}

type CallbackEventArgs struct {
	Method     string
	Parameters interface{}
}
