package other

import "math"

// 整数反转
// https://leetcode.cn/problems/reverse-integer/
func reverse(x int) int {
	ret := 0

	for x != 0 {
		if ret < math.MinInt32/10 || ret > math.MaxInt32/10 {
			return 0
		}

		digit := x % 10
		x = x / 10

		ret = ret*10 + digit
	}

	return ret
}
