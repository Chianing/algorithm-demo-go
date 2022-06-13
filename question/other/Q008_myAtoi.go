package other

import (
	"math"
	"strings"
)

// 字符串转整数
// https://leetcode.cn/problems/string-to-integer-atoi/
func myAtoi(s string) int {
	s = strings.Trim(s, " ")
	if len(s) == 0 {
		return 0
	}

	sign := 1
	ret := 0
	i := 0

	if s[0] == '-' {
		sign = -1
		i++
	} else if s[0] == '+' {
		sign = 1
		i++
	}

	for ; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return sign * ret
		}

		ret = 10*ret + int(s[i]-'0')
		if sign*ret > math.MaxInt32 {
			return math.MaxInt32
		}
		if sign*ret < math.MinInt32 {
			return math.MinInt32
		}
	}

	return sign * ret

}
