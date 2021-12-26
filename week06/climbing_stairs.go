package week06

func climbStairs(n int) int {
	f := make([]int, n+1)
	f[0], f[1] = 1, 1
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}

// 直接可以用变量，内存更少
// func climbStairs(n int) int {
//     prePre ,pre,ans := 1,1,1
//     for i:=2;i<=n;i++{
//         ans = pre + prePre
//         pre , prePre = ans , pre
//     }
//     return ans
// }
