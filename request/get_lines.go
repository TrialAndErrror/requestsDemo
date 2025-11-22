package request

import "strings"

func getRequestLines(requestText string) []string {
	// Clean carriage returns and replace with simple newlines
	normalizedText := strings.ReplaceAll(requestText, "\r\n", "\n")

	// Split request into slice of strings
	cleanLines := strings.Split(normalizedText, "\n")
	return cleanLines
}
