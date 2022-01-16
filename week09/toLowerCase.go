package week09

func toLowerCase(s string) string {
	//return strings.ToLower(s)
	b := []byte(s)
	for i := 0; i < len(s); i++ {
		if b[i] >= 'A' && b[i] <= 'Z' {
			b[i] = b[i] - 'A' + 'a'
		}
	}
	return string(b)
}
