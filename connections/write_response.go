package connections

import (
	"log"
	"net"
)

func writeResponse(conn net.Conn, response string) error {
	_, err := conn.Write([]byte(response))
	if err != nil {
		log.Printf("Error writing to connection: %v", err)
		return err
	}
	return nil
}
