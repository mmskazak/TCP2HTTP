package isrequest

// IsRequestFunc defines a function type that determines whether a message is a request.
type IsRequestFunc func([]byte) bool
