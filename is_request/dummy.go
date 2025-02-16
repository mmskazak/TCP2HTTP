package isrequest

// Dummy is a default request checker that always returns false
func IsDummy(message []byte) bool {
	return false
}
