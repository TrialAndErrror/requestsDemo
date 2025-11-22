package request

import (
	"encoding/json"
	"log"
	"strings"
)

type RawRequest struct {
	Meta Meta
	Data []string
}

func GetRawRequest(meta string, requestData []string) RawRequest {
	requestMetaParts := strings.Split(meta, " ")
	return RawRequest{
		Meta: Meta{
			method:   requestMetaParts[0],
			endpoint: requestMetaParts[1],
			version:  requestMetaParts[2],
		},
		Data: requestData,
	}
}

func GetRequestLines(requestText string) []string {
	// Clean carriage returns and replace with simple newlines
	normalizedText := strings.ReplaceAll(requestText, "\r\n", "\n")

	// Split request into slice of strings
	cleanLines := strings.Split(normalizedText, "\n")
	return cleanLines
}

type Meta struct {
	method   string
	endpoint string
	version  string
}

type ProcessedRequest struct {
	Meta    Meta
	Headers map[string]string

	// GET requests have Params
	Params map[string]interface{}

	// POST requests have Body and Data
	Body string
	Data map[string]interface{}
}

func parseKeyValuePairs(params string) map[string]interface{} {
	paramsMap := make(map[string]interface{})
	paramParts := strings.Split(params, "&")
	for i := 0; i < len(paramParts); i++ {
		keyValue := strings.SplitN(paramParts[i], "=", 2)
		key, value := keyValue[0], keyValue[1]
		paramsMap[key] = value
	}
	return paramsMap
}

func parseQueryParams(endpoint string) map[string]interface{} {
	endpointPartsList := strings.SplitN(endpoint, "?", 2)
	if len(endpointPartsList) < 2 {
		log.Printf("Error parsing query params: %v", endpointPartsList)
		return map[string]interface{}{}
	}
	paramsString := endpointPartsList[1]
	return parseKeyValuePairs(paramsString)
}

func processGetRequest(request RawRequest) ProcessedRequest {
	// Separate Headers from Body and Data
	headerCount := len(request.Data) - 1
	headers := request.Data[:headerCount]

	// Parse Headers
	headersMap := parseHeaders(headers)

	// Parse query params
	paramsMap := parseQueryParams(request.Meta.endpoint)

	return ProcessedRequest{
		Meta:    request.Meta,
		Params:  paramsMap,
		Headers: headersMap,
	}
}

func parseHeaders(headers []string) map[string]string {
	headersMap := make(map[string]string)
	for i := 0; i < len(headers); i++ {
		// Skip blank lines that are processed as headers
		if len(headers[i]) == 0 {
			continue
		}
		parts := strings.SplitN(headers[i], ":", 2)
		if len(parts) != 2 {
			log.Println("Invalid header:", headers[i])
			continue
		}
		headersMap[parts[0]] = parts[1]
	}

	return headersMap
}

func processPostRequest(request RawRequest) ProcessedRequest {
	log.Printf("Processing request: %+v", request)
	// Separate Headers from Body and Data
	headerCount := len(request.Data) - 2
	headers := request.Data[:headerCount]
	body := request.Data[headerCount]
	data := request.Data[headerCount+1]

	// Parse Headers
	headersMap := parseHeaders(headers)

	if len(data) == 0 {
		return ProcessedRequest{
			Meta:    request.Meta,
			Headers: headersMap,
			Body:    body,
		}
	}

	var contentType string
	for key, value := range headersMap {
		if key == "Content-Type" {
			contentType = value
		}
	}

	switch contentType {
	case "application/json":
		jsonData := make(map[string]interface{})
		err := json.Unmarshal([]byte(data), &jsonData)
		if err != nil {
			log.Println("Failed to parse request: ", err)
			return ProcessedRequest{}
		}

		return ProcessedRequest{
			Meta:    request.Meta,
			Headers: headersMap,
			Body:    body,
			Data:    jsonData,
		}
	case "x-www-form-urlencoded":
		formData := parseKeyValuePairs(data)
		return ProcessedRequest{
			Meta:    request.Meta,
			Headers: headersMap,
			Body:    body,
			Data:    formData,
		}
	}

	return ProcessedRequest{
		Meta:    request.Meta,
		Headers: headersMap,
		Body:    body,
	}
}

func ProcessRequest(requestText string) ProcessedRequest {
	requestLines := GetRequestLines(requestText)
	rawRequest := GetRawRequest(requestLines[0], requestLines[1:])
	switch rawRequest.Meta.method {
	case "GET":
		return processGetRequest(rawRequest)
	case "POST":
		return processPostRequest(rawRequest)
	}
	return ProcessedRequest{}
}
