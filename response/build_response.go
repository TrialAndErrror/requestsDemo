package response

import (
	"fmt"
	"strconv"
	"time"
)

func buildResponse(status string, responseBody string) string {
	headers := map[string]string{
		"Date":           fmt.Sprintf("%v", time.Now().Format(time.RFC1123)),
		"Server":         "Handmade Golang Server 1.0",
		"Content-Length": strconv.Itoa(len(responseBody)),
		"Content-Type":   "text/html",
	}

	headersString := ""
	for key, value := range headers {
		headersString += fmt.Sprintf("%s: %s\n", key, value)
	}

	return fmt.Sprintf("%s\n%s\n\n%s", status, headersString, responseBody)
}
