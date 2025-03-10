package structs

type TMModeScriptInfo struct {
	Name               string                 `json:"Name" xmlrpc:"Name"`
	CompatibleMapTypes string                 `json:"CompatibleMapTypes" xmlrpc:"CompatibleMapTypes"`
	Description        string                 `json:"Description" xmlrpc:"Description"`
	Version            string                 `json:"Version" xmlrpc:"Version"`
	ParamDescs         []TMScriptParamDescs   `json:"ParamDescs" xmlrpc:"ParamDescs"`
	CommandDescs       []TMScriptCommandDescs `json:"CommandDescs" xmlrpc:"CommandDescs"`
}

type TMScriptParamDescs struct {
	Name    string `json:"Name" xmlrpc:"Name"`
	Desc    string `json:"Desc" xmlrpc:"Desc"`
	Type    string `json:"Type" xmlrpc:"Type"`
	Default string `json:"Default" xmlrpc:"Default"`
}

type TMScriptCommandDescs struct {
	Name    string `json:"Name" xmlrpc:"Name"`
	Desc    string `json:"Desc" xmlrpc:"Desc"`
	Type    string `json:"Type" xmlrpc:"Type"`
	Default string `json:"Default" xmlrpc:"Default"`
}

type TMModeScriptSettings map[string]interface{}
