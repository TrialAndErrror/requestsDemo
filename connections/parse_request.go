package connections

import (
	"fmt"
	"net"
)

func parseRequest(conn net.Conn, requestBuffer []byte) (interface{}, error) {
	fmt.Printf("Received from %s: %s\n", conn.RemoteAddr(), string(requestBuffer))
	return nil, nil
}
