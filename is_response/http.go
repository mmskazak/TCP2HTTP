package isresponse

import "strings"

func IsHTTPResponse(message []byte) bool {
	return strings.HasPrefix(string(message), "HTTP/")
}
