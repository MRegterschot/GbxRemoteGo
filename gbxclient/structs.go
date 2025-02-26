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

type GbxCallbacks struct {
	OnAnyCallback []func(*GbxClient, CallbackEventArgs)
	OnBeginMap []func(*GbxClient, events.MapEventArgs)
	OnBeginMatch []func(*GbxClient)
	OnEcho []func(*GbxClient, events.EchoEventArgs)
	OnEndMap []func(*GbxClient, events.MapEventArgs)
	OnEndMatch []func(*GbxClient, events.EndMatchEventArgs)
	OnMapListModified []func(*GbxClient, events.MapListModifiedEventArgs)
	OnPlayerAlliesChanged []func(*GbxClient, events.PlayerAlliesChangedEventArgs)
	OnPlayerChat []func(*GbxClient, events.PlayerChatEventArgs)
	OnPlayerConnect []func(*GbxClient, events.PlayerConnectEventArgs)
	OnPlayerDisconnect []func(*GbxClient, events.PlayerDisconnectEventArgs)
	OnPlayerInfoChanged []func(*GbxClient, events.PlayerInfoChangedEventArgs)
	OnPlayerManialinkPageAnswer []func(*GbxClient, events.PlayerManialinkPageAnswerEventArgs)
	OnServerStart []func(*GbxClient)
	OnServerStop []func(*GbxClient)
	OnStatusChanged []func(*GbxClient, events.StatusChangedEventArgs)
	OnTunnelDataReceived []func(*GbxClient, events.TunnelDataReceivedEventArgs)
	OnVoteUpdated []func(*GbxClient, events.VoteUpdatedEventArgs)
	OnPlayerCheckpoint []func(*GbxClient, events.PlayerWayPointEventArgs)
	OnPlayerFinish []func(*GbxClient, events.PlayerWayPointEventArgs)
	OnPlayerIncoherence []func(*GbxClient, events.PlayerIncoherenceEventArgs)
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
