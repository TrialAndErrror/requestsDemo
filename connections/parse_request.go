package connections

import (
	"log"
	"net"
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
