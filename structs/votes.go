package structs

type TMCurrentCallVote struct {
	CallerLogin string `json:"CallerLogin" xmlrpc:"CallerLogin"`
	CmdName     string `json:"CmdName" xmlrpc:"CmdName"`
	CmdParam    string `json:"CmdParam" xmlrpc:"CmdParam"`
}
