package week03

import "sort"

func permuteUnique(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	n := len(nums)
	ans := [][]int{}
	used := make([]bool, n)
	combine := []int{}
	var recurion func(int)
	recurion = func(pos int) {
		if pos == n {
			ans = append(ans, append([]int(nil), combine...))
			return
		}
		for i := 0; i < n; i++ {
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}
			if !used[i] {
				combine = append(combine, nums[i])
				used[i] = true
				recurion(pos + 1)
				used[i] = false
				combine = combine[:len(combine)-1]
			}
		}
	}
	recurion(0)
	return ans
}
