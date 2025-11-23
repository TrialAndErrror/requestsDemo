package request

import "strings"

func getRawRequest(meta string, requestData []string) RawRequest {
	requestMetaParts := strings.Split(meta, " ")
	return RawRequest{
		Meta: Meta{
			Method:   requestMetaParts[0],
			Endpoint: requestMetaParts[1],
			Version:  requestMetaParts[2],
		},
		Data: requestData,
	}
}
