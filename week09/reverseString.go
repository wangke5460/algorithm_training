package week09

func reverseString(s []byte) {
	for l, r := 0, len(s)-1; l < r; l++ {
		s[l], s[r] = s[r], s[l]
		r--
	}
}
