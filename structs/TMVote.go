package structs

type TMCurrentCallVote struct {
	CallerLogin string `json:"CallerLogin"`
	CmdName     string `json:"CmdName"`
	CmdParam    string `json:"CmdParam"`
}
