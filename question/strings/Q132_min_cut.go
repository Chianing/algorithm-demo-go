package strings

// 分隔回文串 2
// https://leetcode.cn/problems/palindrome-partitioning-ii/
func minCut(s string) int {
	length := len(s)
	strArr := []byte(s)
	// 记录每个长度下 最小的回文子串个数
	retArr := make([]int, length+1)

	// 长度为0时 回文子串个数为0 所以从1开始计算
	for i := 1; i <= length; i++ {
		retArr[i] = i
		for j := 0; j < i; j++ {
			if isReverse(strArr, j, i-1) {
				// 新增字符串后 可以构成新增的子串
				retArr[i] = minNum(retArr[i], retArr[j]+1)
			}
		}
	}

	return retArr[length] - 1
}

func isReverse(strArr []byte, left, right int) bool {
	for left < right {
		if strArr[left] != strArr[right] {
			return false
		}
		left++
		right--
	}

	return true
}

func minNum(a, b int) int {
	if a > b {
		return b
	}
	return a
}
