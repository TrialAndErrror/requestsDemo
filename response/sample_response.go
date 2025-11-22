package response

import (
	"bytes"
	"html/template"
	"log"
)

func MakeSampleResponse() (string, error) {
	responseTemplate, err := template.ParseFiles("templates/sample-response.html")
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		return "", err
	}

	var buf bytes.Buffer
	err = responseTemplate.Execute(&buf, nil)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		return "", err
	}

	response := buildResponse("HTTP/1.1 200 OK", buf.String())
	return response, nil
}

func MakeGenericErrorResponse() string {
	statusCode := "HTTP/1.1 500 Internal Server Error"
	responseBody := `<!DOCTYPE html><html lang="en"><head><title>Error: Internal Server Error</title></head><body><h1>An unspecified server error occurred. Please try again</h1></body></html>`
	return buildResponse(statusCode, responseBody)
}
