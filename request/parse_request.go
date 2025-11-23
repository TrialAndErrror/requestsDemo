package request

import (
	"fmt"
	"log"
)

func processGetRequest(request RawRequest) ProcessedRequest {
	// Separate Headers from Body and Data
	headerCount := len(request.Data) - 1
	headers := request.Data[:headerCount]

	// Parse Headers
	headersMap := parseHeaders(headers)

	// Parse query params
	paramsMap := parseQueryParams(request.Meta.Endpoint)

	return ProcessedRequest{
		Meta:    request.Meta,
		Params:  paramsMap,
		Headers: headersMap,
	}
}

func processPostRequest(request RawRequest) ProcessedRequest {
	// Separate Headers from Body and Data
	headerCount := len(request.Data) - 2
	headers := request.Data[:headerCount]

	// Skip request.Data[headerCount] as it is a blank line

	data := request.Data[headerCount+1]

	// Parse Headers
	headersMap := parseHeaders(headers)

	processedRequest := ProcessedRequest{
		Meta:    request.Meta,
		Headers: headersMap,
	}

	// If no form or JSON data to parse, return the request where we're at
	if len(data) == 0 {
		return processedRequest
	}

	contentTypeRaw, ok := headersMap["Content-Type"]
	if !ok {
		log.Printf("Missing Content-Type header: %v", headers)
		return processedRequest
	}

	contentType := GetContentType(contentTypeRaw)

	switch contentType {
	case "application/json":
		jsonData := parseJSONData(data)
		processedRequest.Data = jsonData

	case "application/x-www-form-urlencoded":
		formData := parseKeyValuePairs(data)
		processedRequest.Data = formData
	}
	return processedRequest
}

func ProcessRequest(requestText string) (ProcessedRequest, error) {
	requestLines := getRequestLines(requestText)
	rawRequest := getRawRequest(requestLines[0], requestLines[1:])
	switch rawRequest.Meta.Method {
	case "GET":
		return processGetRequest(rawRequest), nil
	case "POST":
		return processPostRequest(rawRequest), nil
	}
	return ProcessedRequest{}, fmt.Errorf("failed to process request; invalid method %s", rawRequest.Meta.Method)
}
