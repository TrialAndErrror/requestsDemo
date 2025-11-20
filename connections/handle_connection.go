package connections

import (
	"fmt"
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

func sendGenericResponse(conn net.Conn) error {
	statusCode := "HTTP/1.1 200 OK"

	responseBody := `<!DOCTYPE html><html><head><title>Example</title></head><body><h1>Hello,World!</h1></body></html>`

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

	responseData := fmt.Sprintf("%s\n%s\n\n%s", statusCode, headersString, responseBody)

	response := []byte(responseData)

	_, err := conn.Write(response)
	return err
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

	err = sendGenericResponse(conn)
	if err != nil {
		log.Printf("Error writing to connection: %v", err)
		return
	}
}
