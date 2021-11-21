package week01

func maximalRectangle(matrix [][]byte) int {
	ans := 0
	for i, rowsVal := range matrix {
		eveRectHeihts := make([]int, 0, len(matrix[0]))
		for j, v := range rowsVal {
			calmulate := 0
			if v == '1' {
				calmulate++
				for n := i - 1; n >= 0; n-- {
					if matrix[n][j] == '1' {
						calmulate++
					} else {
						break
					}
				}
			}
			eveRectHeihts = append(eveRectHeihts, calmulate)
		}
		ans = max(ans, largestRectangleArea(eveRectHeihts))
	}
	return ans
}

func largestRectangleArea(heights []int) int {
	//使栈弹空
	heights = append(heights, 0)
	//维护一个单调栈
	type rect struct {
		width  int
		height int
	}
	s := []rect{}
	ans := 0

	for _, height := range heights {
		accumulatedWidth := 0
		//栈顶之前的高度 >= 当前高度，单调性破坏，确定了栈顶高度的扩张范围，需要删除栈顶
		for len(s) != 0 && s[len(s)-1].height >= height {
			accumulatedWidth += s[len(s)-1].width
			ans = max(ans, s[len(s)-1].height*accumulatedWidth)
			s = s[:len(s)-1]
		}
		s = append(s, rect{width: accumulatedWidth + 1, height: height})
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
