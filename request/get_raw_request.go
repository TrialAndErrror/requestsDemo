package request

import "strings"

func GetRawRequest(meta string, requestData []string) RawRequest {
	requestMetaParts := strings.Split(meta, " ")
	return RawRequest{
		Meta: Meta{
			method:   requestMetaParts[0],
			endpoint: requestMetaParts[1],
			version:  requestMetaParts[2],
		},
		Data: requestData,
	}
}
