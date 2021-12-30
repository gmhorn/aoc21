package lib

// IntAbs returns the absolute value for the integer argument.
//
// Maybe there's a faster way, this is Good Enough for now.
func IntAbs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
