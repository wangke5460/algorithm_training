package week06

//f[i] 位置为i时最长子序列长度  j< i时,nums[j] < nums[i]才能f[i] = max(f[j]+1)
//g[i] 位置为i时最长子序列个数 f[i] > f[j] + 1，g[i] = g[j],f[i] == f[j] + 1时，g[i] += g[j]
func findNumberOfLIS(nums []int) int {
	n := len(nums)
	f := make([]int, n)
	g := make([]int, n)
	f[0] = 1
	g[0] = 1
	for i := 1; i < n; i++ {
		f[i] = 1
		g[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				if f[i] < f[j]+1 {
					f[i] = f[j] + 1
					g[i] = g[j]
					continue
				}
				if f[i] == f[j]+1 {
					g[i] += g[j]
					continue
				}
			}
		}
	}

	longSub := f[0]
	ans := g[0]
	for i := 1; i < n; i++ {
		if longSub < f[i] {
			longSub = f[i]
			ans = g[i]
			continue
		}
		if longSub == f[i] {
			ans += g[i]
			continue
		}
	}
	return ans
}
