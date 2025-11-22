package connections

import (
	"bytes"
	"html/template"
	"log"
	"net"
	"requestsDemo/request"
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

	requestString := string(requestBytes)
	requestData := request.ProcessRequest(requestString)
	if err != nil {
		return err
	}
	log.Printf("Received request %+v", requestData)

	responseTemplate, err := template.ParseFiles("templates/sample-response.html")
	if err != nil {
		return writeResponse(conn, buildGenericErrorResponse())
	}

	var buf bytes.Buffer
	err = responseTemplate.Execute(&buf, nil)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		return writeResponse(conn, buildGenericErrorResponse())
	}

	response := buildResponse("HTTP/1.1 200 OK", buf.String())
	return writeResponse(conn, response)
}
