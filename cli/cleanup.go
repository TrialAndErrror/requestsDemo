package cli

import (
	"fmt"
	"net"
	"os"
)

func CleanupListener(l net.Listener) {
	err := l.Close()
	if err != nil {
		fmt.Printf("error closing listener: %s", err)
		os.Exit(1)
	}
}
