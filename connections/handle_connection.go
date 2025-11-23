package connections

import (
	"log"
	"net"
	"requestsDemo/response"
)

func cleanupConnection(conn net.Conn) {
	err := conn.Close()
	if err != nil {
		log.Printf("Error closing connection: %v", err)
		return
	}
}

func HandleConnection(conn net.Conn) error {
	defer cleanupConnection(conn)

	responseBody := "<!DOCTYPE html><html><head><meta charset=\"UTF-8\"></head><body>Hello, World!</body></html>"
	responseString := response.BuildResponse("HTTP/1.1 200 OK", responseBody)

	return writeResponse(conn, responseString)
}
