package connections

import (
	"log"
	"net"
	"requestsDemo/request"
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
	requestData, err := request.ProcessRequest(requestString)
	if err != nil {
		log.Printf("Error parsing request: %v", err)
		return writeResponse(conn, response.MakeGenericErrorResponse())
	}

	log.Printf("Received request %+v", requestData)

	responseString, err := response.MakeResponse(requestData)
	if err != nil {
		log.Printf("Error generating response: %v", err)
		return writeResponse(conn, response.MakeGenericErrorResponse())
	}

	return writeResponse(conn, responseString)
}
