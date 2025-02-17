package main

import (
	"errors"
)

type TMMethodSignature struct {
	ReturnType string
	ParamTypes []string
}

// Given the name of a method, return an array of legal signatures. Each signature is an array of strings. The first item of each signature is the return type, and any others items are parameter types.
func (client *GbxClient) MethodSignature(signature string) ([]TMMethodSignature, error) {
	res, err := client.Call("system.methodSignature", signature)
	if err != nil {
		return nil, err
	}

	// Ensure the response is a slice
	data, ok := res.([]interface{})
	if !ok {
		return nil, errors.New("unexpected response format")
	}

	// Convert slice to []TMMethodSignature
	methods := make([]TMMethodSignature, len(data))
	for i, v := range data {
		method := TMMethodSignature{}
		if v == nil {
			method.ReturnType = ""
			method.ParamTypes = nil
		} else {
			signature := v.([]interface{})
			method.ReturnType = signature[0].(string)

			if len(signature) > 1 {
				paramTypes := make([]string, len(signature)-1)
				for j := 1; j < len(signature); j++ {
					paramTypes[j-1] = signature[j].(string)
				}
				method.ParamTypes = paramTypes
			} else {
				method.ParamTypes = nil
			}
		}

		methods[i] = method
	}

	return methods, nil
}