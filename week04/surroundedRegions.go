package week04

func solve(board [][]byte) {
	//可以找出在边界上的所有连通块
	m := len(board)
	n := len(board[0])
	//判断是否在边界上
	isBoundary := func(x, y int) bool {
		return x == 0 || y == 0 || x == m-1 || y == n-1
	}
	isBoundaryValue := false

	//判断下标是否合法
	indexValid := func(x, y int) bool {
		return x >= 0 && y >= 0 && x < m && y < n
	}
	//每个点访问一次
	visied := make([][]bool, m)
	for i := 0; i < m; i++ {
		visied[i] = make([]bool, n)
	}
	//方向数组
	dx := []int{-1, 0, 1, 0}
	dy := []int{0, 1, 0, -1}

	//临时变量
	tempResult := [][]int{}

	//寻找连通块
	var dfs func(int, int)
	dfs = func(x, y int) {
		visied[x][y] = true
		if isBoundary(x, y) {
			isBoundaryValue = true
		}
		tempResult = append(tempResult, []int{x, y})
		for k := 0; k < 4; k++ {
			nx := x + dx[k]
			ny := y + dy[k]
			if !indexValid(nx, ny) || board[nx][ny] == 'X' || visied[nx][ny] {
				continue
			}
			dfs(nx, ny)
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' && !visied[i][j] {
				dfs(i, j)
				if !isBoundaryValue {
					for _, res := range tempResult {
						board[res[0]][res[1]] = 'X'
					}
				}
				isBoundaryValue = false
				tempResult = [][]int{}
			}
		}
	}
}
