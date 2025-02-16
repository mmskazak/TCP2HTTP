package middleware

import (
	"fmt"

	"github.com/mmskazak/tcpwrapper"
)

// LogMiddleware creates a middleware that logs the message length and first N bytes
func LogMiddleware(prefix string, maxPreviewBytes int) tcpwrapper.Middleware {
	return func(data []byte) ([]byte, error) {
		previewLen := len(data)
		if previewLen > maxPreviewBytes {
			previewLen = maxPreviewBytes
		}
		fmt.Printf("[%s] Message length: %d bytes, Preview: %s\n",
			prefix,
			len(data),
			string(data[:previewLen]))
		return data, nil
	}
}
