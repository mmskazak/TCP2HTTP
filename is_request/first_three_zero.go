package isrequest

// IsFirstThreeZero checks if the first three bytes of the message are zero.
func IsFirstThreeZero(message []byte) bool {
	return message[0] == 0 && message[1] == 0 && message[2] == 0
}
