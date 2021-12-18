package week05

func shipWithinDays(weights []int, days int) int {
	left, right := 0, 0
	for _, weight := range weights {
		left = max(left, weight)
		right += weight
	}
	for left < right {
		mid := (left + right) / 2
		if validday(weights, days, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}

//每个船装载为size，能否5天装完
func validday(weights []int, days, size int) bool {
	ship := 0
	realday := 1
	for _, weight := range weights {
		//if weight > size {return false}
		if ship+weight <= size {
			ship += weight
		} else {
			realday++
			ship = weight
		}
	}
	return realday <= days
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
