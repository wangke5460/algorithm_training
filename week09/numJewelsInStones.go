package week09

func numJewelsInStones(jewels string, stones string) int {
	hashMap := map[byte]int{}
	for i := 0; i < len(stones); i++ {
		hashMap[stones[i]]++
	}
	count := 0
	for i := 0; i < len(jewels); i++ {
		count += hashMap[jewels[i]]
	}
	return count
}
