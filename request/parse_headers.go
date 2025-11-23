package request

import (
	"log"
	"strings"
)

func parseHeaders(headers []string) map[string]string {
	headersMap := make(map[string]string)
	for i := 0; i < len(headers); i++ {
		// Skip blank lines that are processed as headers
		if len(headers[i]) == 0 {
			continue
		}
		parts := strings.SplitN(headers[i], ":", 2)
		if len(parts) != 2 {
			log.Println("Invalid header:", headers[i])
			continue
		}
		headersMap[parts[0]] = parts[1]
	}

	return headersMap
}
