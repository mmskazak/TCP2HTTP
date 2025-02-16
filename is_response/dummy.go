package isresponse

//Response for all messages
func IsDummy(message []byte) bool {
	return true
}

// Dummy is a default response checker that always returns false
func Dummy(message []byte) bool {
	return false
}
