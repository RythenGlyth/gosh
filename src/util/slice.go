package util

// ContainsR checks, whether x is in rs and returns the position.
// If it is not, -1 is returned.
func ContainsR(x rune, rs []rune) int {
	for i, r := range rs {
		if r == x {
			return i
		}
	}

	return -1
}

func ContainsS(x string, ss []string) int {
	for i, s := range ss {
		if s == x {
			return i
		}
	}

	return -1
}
