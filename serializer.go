package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"
)

type MethodCall struct {
	MethodName string  `xml:"methodName"`
	Params     []Param `xml:"params>param"`
}

type MethodResponse struct {
	Params []Param `xml:"params>param"`
	Fault  *Fault  `xml:"fault"`
}

type Fault struct {
	FaultCode   *int    `xml:"value>struct>member>value>int"`
	FaultString *string `xml:"value>struct>member>value>string"`
}

type Param struct {
	Value Value `xml:"value"`
}

type Value struct {
	String   *string  `xml:"string"`
	Int      *int     `xml:"i4"`
	Bool     *bool    `xml:"boolean"`
	Float    *float64 `xml:"double"`
	Struct   *Struct  `xml:"struct"`
	Array    *Array   `xml:"array"`
	Base64   *string  `xml:"base64"`
	DateTime *string  `xml:"dateTime.iso8601"`
}

type Struct struct {
	Members []Member `xml:"member"`
}

type Member struct {
	Name  string `xml:"name"`
	Value Value  `xml:"value"`
}

type Array struct {
	Data []Value `xml:"data>value"`
}

type CData string

type XMLParam struct {
	Value string `xml:",innerxml"`
}

type XMLRequest struct {
	XMLName    xml.Name   `xml:"methodCall"`
	MethodName string     `xml:"methodName"`
	Params     []XMLParam `xml:"params>param"`
}

func DeserializeMethodResponse(data []byte) (interface{}, error) {
	sanitizedData := sanitizeXML(string(data))
	if err := checkResponse([]byte(sanitizedData)); err != nil {
		return nil, err
	}
	r := bytes.NewReader([]byte(sanitizedData))

	var methodResponse MethodResponse
	decoder := xml.NewDecoder(r)

	// Parse the XML into the methodResponse structure
	err := decoder.Decode(&methodResponse)
	if err != nil {
		return nil, err
	}

	// If fault is present, handle it
	if methodResponse.Fault != nil {
		if methodResponse.Fault.FaultCode != nil && methodResponse.Fault.FaultString != nil {
			return nil, fmt.Errorf("FaultCode: %d, FaultString: %s", *methodResponse.Fault.FaultCode, *methodResponse.Fault.FaultString)
		}
		return nil, errors.New("fault found but missing details")
	}

	// Check if there are parameters, if not return an error
	if len(methodResponse.Params) == 0 {
		return nil, errors.New("no parameters found")
	}

	param := methodResponse.Params[0]
	return deserializeValue(param.Value)
}

func DeserializeMethodCall(data []byte) (string, interface{}, error) {
	sanitizedData := sanitizeXML(string(data))
	if err := checkResponse([]byte(sanitizedData)); err != nil {
		return "", nil, err
	}
	r := bytes.NewReader([]byte(sanitizedData))

	var methodCall MethodCall
	decoder := xml.NewDecoder(r)

	// Parse the XML into the methodCall structure
	err := decoder.Decode(&methodCall)
	if err != nil {
		return "", nil, err
	}

	// If there are no parameters, return the method name and nil for params
	if len(methodCall.Params) == 0 {
		return methodCall.MethodName, nil, nil
	}

	// Parse the parameters
	params := make([]interface{}, len(methodCall.Params))
	for i, param := range methodCall.Params {
		value, err := deserializeValue(param.Value)
		if err != nil {
			return "", nil, err
		}
		params[i] = value
	}

	return methodCall.MethodName, params, nil
}

func xmlSerializer(method string, params []interface{}) (string, error) {
	var xmlParams []XMLParam
	for _, param := range params {
		// Use reflection to handle different types of params
		paramStr, err := serializeParam(param)
		if err != nil {
			return "", err
		}
		xmlParams = append(xmlParams, XMLParam{Value: paramStr})
	}

	request := XMLRequest{
		MethodName: method,
		Params:     xmlParams,
	}

	requestData, err := xml.Marshal(request)
	if err != nil {
		return "", err
	}

	return `<?xml version="1.0"?>` + string(requestData), nil
}

// Helper function to serialize different types
func serializeParam(param interface{}) (string, error) {
	// Handle pointers by dereferencing them if not nil
	val := reflect.ValueOf(param)
	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return "", nil // Represent nil pointers as <nil/>
		}
		return serializeParam(val.Elem().Interface()) // Dereference and reprocess
	}

	switch v := param.(type) {
	case string:
		return fmt.Sprintf("<value><string>%s</string></value>", v), nil
	case int, int32, int64:
		return fmt.Sprintf("<value><int>%d</int></value>", v), nil
	case float32, float64:
		return fmt.Sprintf("<value><double>%f</double></value>", v), nil
	case bool:
		if v {
			return "<value><boolean>1</boolean></value>", nil
		}
		return "<value><boolean>0</boolean></value>", nil
	case []interface{}: // Handle arrays (slice of values)
		var values string
		for _, elem := range v {
			serializedElem, err := serializeParam(elem)
			if err != nil {
				return "", err
			}
			values += fmt.Sprintf("%s", serializedElem)
		}
		return fmt.Sprintf("<value><array><data>%s</data></array></value>", values), nil
	case []byte: // Handle base64 encoding
		encoded := base64.StdEncoding.EncodeToString(v)
		return fmt.Sprintf("<value><base64>%s</base64></value>", encoded), nil
	case time.Time: // Handle date/time serialization
		return fmt.Sprintf("<value><dateTime.iso8601>%s</dateTime.iso8601></value>", v.Format("20060102T15:04:05Z")), nil
	case map[string]interface{}: // Handle struct serialization (map of name-value pairs)
		var members string
		for name, value := range v {
			serializedValue, err := serializeParam(value)
			if err != nil {
				return "", err
			}
			members += fmt.Sprintf("<member><name>%s</name>%s</member>", name, serializedValue)
		}
		return fmt.Sprintf("<value><struct>%s</struct></value>", members), nil
	case CData: // Handle CDATA serialization
		return fmt.Sprintf("<value><string><![CDATA[%s]]></string></value>", v), nil
	case nil: // Handle nil serialization
		return "<value><nil/></value>", nil
	default:
		// Handle custom structs using reflection
		val := reflect.ValueOf(param)
		if val.Kind() == reflect.Struct {
			var members string
			for i := 0; i < val.NumField(); i++ {
				field := val.Type().Field(i)
				fieldInterface := val.Field(i).Interface()
				fieldValue := reflect.ValueOf(fieldInterface)
				if fieldValue.Kind() == reflect.Ptr && fieldValue.IsNil() {
					continue // Skip zero values
				}
				serializedValue, err := serializeParam(fieldInterface)
				if err != nil {
					return "", err
				}
				members += fmt.Sprintf("<member><name>%s</name>%s</member>", field.Name, serializedValue)
			}
			return fmt.Sprintf("<value><struct>%s</struct></value>", members), nil
		}
		// Handle unsupported types explicitly
		return "", fmt.Errorf("unsupported parameter type: %T", param)
	}
}

// Recursive function to handle deserialization of a single value
func deserializeValue(value Value) (interface{}, error) {
	switch {
	case value.String != nil:
		return *value.String, nil
	case value.Int != nil:
		return *value.Int, nil
	case value.Bool != nil:
		return *value.Bool, nil
	case value.Float != nil:
		return *value.Float, nil
	case value.Struct != nil:
		// Handle structs by converting to a map
		parsedData := make(map[string]interface{})
		for _, member := range value.Struct.Members {
			memberValue, err := deserializeValue(member.Value) // Recursively deserialize each member's value
			if err != nil {
				return nil, err
			}
			parsedData[member.Name] = memberValue
		}
		return parsedData, nil
	case value.Array != nil:
		// Handle arrays, process each element in the array
		parsedArray := make([]interface{}, len(value.Array.Data))
		for i, item := range value.Array.Data {
			itemValue, err := deserializeValue(item) // Recursively deserialize array items
			if err != nil {
				return nil, err
			}
			parsedArray[i] = itemValue
		}
		return parsedArray, nil
	case value.Base64 != nil:
		return *value.Base64, nil
	case value.DateTime != nil:
		return *value.DateTime, nil
	default:
		return nil, errors.New("unsupported data type or empty value")
	}
}

// Generic function to convert response to a specific type
func convertToStruct(res interface{}, targetType interface{}) error {
	// Ensure the response is either map[string]interface{} or []interface{}
	switch v := res.(type) {
	case map[string]interface{}:
		// Convert map to JSON
		jsonData, err := json.Marshal(v)
		if err != nil {
			return err
		}

		// Ensure targetType is a pointer
		targetVal := reflect.ValueOf(targetType)
		if targetVal.Kind() != reflect.Ptr || targetVal.IsNil() {
			return errors.New("target type must be a non-nil pointer")
		}

		// Unmarshal JSON into the target struct
		err = json.Unmarshal(jsonData, targetType)
		if err != nil {
			return err
		}

	case []interface{}:
		// Handle slice of interfaces
		// Convert the slice to JSON
		jsonData, err := json.Marshal(v)
		if err != nil {
			return err
		}

		// Ensure targetType is a pointer to a slice (e.g., *[]MyStruct)
		targetVal := reflect.ValueOf(targetType)
		if targetVal.Kind() != reflect.Ptr || targetVal.IsNil() {
			return errors.New("target type must be a non-nil pointer")
		}

		// Unmarshal JSON into the target slice
		err = json.Unmarshal(jsonData, targetType)
		if err != nil {
			return err
		}

	default:
		return errors.New("unexpected response format")
	}

	return nil
}

// Sanitize XML input by removing invalid characters
func sanitizeXML(input string) string {
	return strings.Map(func(r rune) rune {
		if r >= 0x20 {
			return r
		}
		return -1 // Remove invalid characters
	}, input)
}

func checkResponse(data []byte) error {
	if bytes.HasSuffix(data, []byte("</methodResponse>")) || bytes.HasSuffix(data, []byte("</methodCall>")) {
		return nil
	}
	return errors.New("incomplete response received")
}
