package isresponse

// Dummy is a default response checker that always returns false
func IsDummy(message []byte) bool {
	return false
}
