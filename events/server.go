package events

type EchoEventArgs struct {
	Internal string `json:"Internal" xmlrpc:"Internal"`
	Public   string `json:"Public" xmlrpc:"Public"`
}

type StatusChangedEventArgs struct {
	StatusCode int    `json:"StatusCode" xmlrpc:"StatusCode"`
	StatusName string `json:"StatusName" xmlrpc:"StatusName"`
}

type TunnelDataReceivedEventArgs struct {
	Login     string `json:"Login" xmlrpc:"Login"`
	PlayerUid int    `json:"PlayerUid" xmlrpc:"PlayerUid"`
	Data      []byte `json:"Data" xmlrpc:"Data"`
}
