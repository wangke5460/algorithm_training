package week05

func minEatingSpeed(piles []int, h int) int {
	left, right := 1, 0
	for _, pile := range piles {
		right = max(right, pile)
	}
	for left < right {
		//mid := (left + right) / 2
		mid := left + (right-left)/2
		if isEaten(piles, h, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}

//一小时能吃k个，h 个 小时 能否吃完
func isEaten(piles []int, h, k int) bool {
	hour := 0
	for _, pile := range piles {
		hour += (pile + k - 1) / k
	}
	return hour <= h
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
