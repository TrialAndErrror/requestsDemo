package cli

import (
	"fmt"
	"os"
)

func ParseArgs(args []string) (port string, parseFailed bool) {
	commandArgs := args[1:]
	if len(commandArgs) < 1 {
		fmt.Printf("Error: No port provided")
		fmt.Printf("Usage: requestsDemo <port>\n")
		return "", true
	}
	return os.Args[1], false
}
