package week02

func subarraySum(nums []int, k int) int {
	//前缀和s[r] - [l]= k
	n := len(nums)
	s := make([]int, n+1)
	for i := 1; i <= n; i++ {
		s[i] = s[i-1] + nums[i-1]
	}
	//两数之差 s[r] - s[l]= k --> s[l] = s[r] - k
	ans := 0
	hashMap := map[int]int{}
	for i := 0; i <= n; i++ {
		//值出现的次数
		_, ok := hashMap[s[i]-k]
		if ok {
			ans = hashMap[s[i]-k] + ans
		}
		hashMap[s[i]]++
	}
	return ans
}
