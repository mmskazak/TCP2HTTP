package isresponse

// IsResponseFunc defines a function type that determines whether a message is a response.
type IsResponseFunc func([]byte) bool
