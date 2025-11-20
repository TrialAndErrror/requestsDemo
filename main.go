package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"requestsDemo/cli"
	"requestsDemo/connections"
)

func main() {
	port, argParseErr := cli.ParseArgs(os.Args)
	if argParseErr {
		os.Exit(1)
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		fmt.Printf("error listening on port %s: %s", port, err)
		os.Exit(1)
	}

	defer cli.CleanupListener(listener)

	fmt.Printf("Listening on port %s\n", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}

		connections.HandleConnection(conn)
	}
}
