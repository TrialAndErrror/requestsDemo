package request

import "strings"

func GetContentType(contentTypeRaw string) string {
	return strings.Trim(strings.SplitN(contentTypeRaw, ";", 2)[0], " ")
}
