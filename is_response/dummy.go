package isresponse

// Dummy is a default response checker that always returns false
func Dummy(message []byte) bool {
	return false
}
