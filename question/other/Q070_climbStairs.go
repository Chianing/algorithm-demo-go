package other

// 爬楼梯
// https://leetcode.cn/problems/climbing-stairs/
func climbStairs(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	ret, pre1, pre2 := 1, 0, 0

	for i := 1; i <= n; i++ {
		pre1 = pre2
		pre2 = ret
		ret = pre1 + pre2
	}

	return ret

}
