package week06

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	f := make([][]int, n)
	f[0] = make([]int, 1)
	f[0][0] = triangle[0][0]
	for i := 1; i < n; i++ {
		f[i] = make([]int, len(triangle[i]))
		for j := 0; j < len(triangle[i]); j++ {
			if j == 0 {
				f[i][j] = f[i-1][j] + triangle[i][j]
				continue
			}
			if j == len(triangle[i])-1 {
				f[i][j] = f[i-1][j-1] + triangle[i][j]
				continue
			}
			f[i][j] = min(f[i-1][j], f[i-1][j-1]) + triangle[i][j]
		}
	}

	ans := f[n-1][0]
	for i := 1; i < len(f[n-1]); i++ {
		ans = min(ans, f[n-1][i])
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

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
