package main

import (
	"bytes"
	"net"
	"sync"
)

type Options struct {
	ShowErrors  bool
	ThrowErrors bool
}

type GbxClient struct {
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

type Callback struct {
	Method string
	Res    interface{}
}

// Deserializer placeholder (you need to implement the actual logic)
type Deserializer struct{}
