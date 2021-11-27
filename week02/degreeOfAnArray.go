package week02

func findShortestSubArray(nums []int) int {
	//要存储三个值：值，出现次数，长度（两坐标）。根据值，更新出现次数和长度
	n := len(nums)
	hashMap := map[int][]int{}
	for i := 0; i < n; i++ {
		_, ok := hashMap[nums[i]]
		//同一个数只需下标最小和最大的两个,中间不用管
		if !ok {
			hashMap[nums[i]] = []int{1, i, i}
		} else {
			hashMap[nums[i]][0]++
			hashMap[nums[i]][2] = i
		}
	}
	//出现频率最大的数字的最小长度,找出hashMap中value最大的几个
	ans := 50001
	frequMax := 0
	for _, v := range hashMap {
		//次数相同，答案取最小的;
		if v[0] == frequMax {
			ans = min(ans, v[2]-v[1]+1)
			frequMax = v[0]
		}
		//答案取 次数大的
		if v[0] > frequMax {
			ans = v[2] - v[1] + 1
			frequMax = v[0]
		}
	}
	return ans

}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
