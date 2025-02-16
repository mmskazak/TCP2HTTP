package isrequest

// Request for all messages
func IsDummy(message []byte) bool {
	return true
}

// Dummy is a default request checker that always returns false
func Dummy(message []byte) bool {
	return false
}
