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

	requestBytes, err := makeRequestBuffer(conn)
	if err != nil {
		log.Printf("Error processing request: %v", err)
		return writeResponse(conn, response.MakeGenericErrorResponse())
	}

	requestString := string(requestBytes)

	log.Printf("Received request: %s", requestString)

	responseString, err := response.MakeSampleResponse()
	if err != nil {
		log.Printf("Error making response: %v", err)
		return writeResponse(conn, response.MakeGenericErrorResponse())
	}

	return writeResponse(conn, responseString)
}
