package structs

type TMMethodSignature struct {
	ReturnType string   `json:"ReturnType" xmlrpc:"ReturnType"`
	ParamTypes []string `json:"ParamTypes" xmlrpc:"ParamTypes"`
}
