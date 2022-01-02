package week07

//f[i]表示i能跳到最远位置
//f[i] = max(i+nums[i],f[i-1])
//只要i不是最后一个，f[i] <= i 则失败
func canJump(nums []int) bool {
	nums = append(make([]int, 1), nums...)
	n := len(nums)
	f := make([]int, n)
	f[0] = 0
	for i := 1; i < n; i++ {
		f[i] = max(f[i-1], i+nums[i])
		if f[i] <= i && i != n-1 {
			return false
		}
	}
	return true
}
func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
