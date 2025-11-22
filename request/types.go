package request

type RawRequest struct {
	Meta Meta
	Data []string
}

type ProcessedRequest struct {
	Meta    Meta
	Headers map[string]string

	// GET requests have Params
	Params map[string]interface{}

	// POST requests have Data
	Data map[string]interface{}
}

type Meta struct {
	method   string
	endpoint string
	version  string
}
