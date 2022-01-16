package week09

func lengthOfLastWord(s string) int {
	index := len(s) - 1
	for s[index] == ' ' {
		index--
	}
	count := 0
	for index >= 0 && s[index] != ' ' {
		count++
		index--
	}
	return count

}
