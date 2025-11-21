package connections

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net"
	"strconv"
	"time"
)

func makeRequestBuffer(conn net.Conn) ([]byte, error) {
	buffer := make([]byte, 1024)
	numBytes, err := conn.Read(buffer)
	if err != nil {
		log.Printf("Error reading from connection: %v", err)
		return buffer, err
	}
	return buffer[:numBytes], nil
}

func buildResponse(status string, responseBody string) string {
	headers := map[string]string{
		"Date":           fmt.Sprintf("%v", time.Now().Format(time.RFC1123)),
		"Server":         "Handmade Golang Server 1.0",
		"Content-Length": strconv.Itoa(len(responseBody)),
		"Content-Type":   "text/html",
	}

	headersString := ""
	for key, value := range headers {
		headersString += fmt.Sprintf("%s: %s\n", key, value)
	}

	return fmt.Sprintf("%s\n%s\n\n%s", status, headersString, responseBody)
}

func buildGenericErrorResponse() string {
	statusCode := "HTTP/1.1 500 Internal Server Error"
	responseBody := `<!DOCTYPE html><html lang="en"><head><title>Error:InternalServerError</title></head><body><h1>An unspecified server error occurred. Please try again</h1></body></html>`
	return buildResponse(statusCode, responseBody)
}

func cleanupConnection(conn net.Conn) {
	err := conn.Close()
	if err != nil {
		log.Printf("Error closing connection: %v", err)
		return
	}
}

func HandleConnection(conn net.Conn) {
	defer cleanupConnection(conn)

	requestBytes, err := makeRequestBuffer(conn)

	_, err = parseRequest(conn, requestBytes)
	if err != nil {
		return
	}

	var response string

	responseTemplate, err := template.ParseFiles("templates/sample-response.html")
	if err != nil {
		response = buildGenericErrorResponse()
	} else {
		var buf bytes.Buffer
		err = responseTemplate.Execute(&buf, nil)
		if err != nil {
			response = buildGenericErrorResponse()
			log.Fatalf("Error executing template: %v", err)
		} else {
			response = buildResponse("HTTP/1.1 200 OK", buf.String())
		}
	}

	_, err = conn.Write([]byte(response))
	if err != nil {
		log.Printf("Error writing to connection: %v", err)
		return
	}
}
