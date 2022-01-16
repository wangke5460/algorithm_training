package week09

func firstUniqChar(s string) int {
	hashMap := map[byte]int{}
	for i := 0; i < len(s); i++ {
		hashMap[s[i]]++
	}
	for i := 0; i < len(s); i++ {
		if hashMap[s[i]] == 1 {
			return i
		}
	}
	return -1
}
