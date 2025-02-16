package isrequest

// Dummy is a default request checker that always returns false
func Dummy(message []byte) bool {
	return false
}
