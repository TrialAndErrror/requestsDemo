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
	Params map[string]string

	// POST requests have Body and Data
	Body string
	Data map[string]interface{}
}

func processGetRequest(request RawRequest) ProcessedRequest {
	// Separate Headers from Body and Data
	headerCount := len(request.Data) - 1
	headers := request.Data[:headerCount]
	body := request.Data[headerCount]

	// Parse Headers
	headersMap := parseHeaders(headers)
	return ProcessedRequest{
		Meta:    request.Meta,
		Params:  map[string]string{},
		Headers: headersMap,
		Body:    body,
	}
}

func parseHeaders(headers []string) map[string]string {
	headersMap := make(map[string]string)
	for i := 0; i < len(headers); i++ {
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
	// Separate Headers from Body and Data
	headerCount := len(request.Data) - 2
	headers := request.Data[:headerCount]
	body := request.Data[headerCount]
	data := request.Data[headerCount+1]

	// Parse Headers
	headersMap := parseHeaders(headers)

	// Parse JSON Data payload
	jsonData := make(map[string]interface{})
	err := json.Unmarshal([]byte(data), &jsonData)
	if err != nil {
		log.Println("Failed to parse request: ", err)
		return ProcessedRequest{}
	}

	return ProcessedRequest{
		Meta:    request.Meta,
		Params:  map[string]string{},
		Headers: headersMap,
		Body:    body,
		Data:    jsonData,
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
