package strings

var str string
var strLen int
var dp [][]bool
var ret [][]string

// 分隔字符串
// https://leetcode.cn/problems/palindrome-partitioning/
func partition(s string) [][]string {
	if len(s) == 0 {
		return [][]string{}
	}

	str = s
	strLen = len(s)
	ret = make([][]string, 0)

	getDPArr(strLen)

	dfs(0, make([]string, 0))

	return ret

}

func dfs(index int, subRet []string) [][]string {
	if index == strLen {
		// 注意 这里subRet底层指向的都是同一个数组
		// 如果使用 ret = append(ret, subRet)会导致最终结果内元素变化
		// 因此这里要么(1)深拷贝一个数组添加到结果集 (2)要么创建一个新的切片
		// (1)
		// subRetCopy := make([]string, len(subRet))
		// copy(subRetCopy, subRet)
		// ret = append(ret, subRetCopy)
		// (2)
		ret = append(ret, append([]string(nil), subRet...))
		return ret
	}

	for i := index; i < strLen; i++ {
		if dp[i][index] {
			subRet = append(subRet, str[index:i+1])
			dfs(i+1, subRet)
			subRet = subRet[:len(subRet)-1]
		}
	}

	return ret
}

func getDPArr(length int) {
	dp = make([][]bool, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]bool, length)
	}

	// 获取所有子串是否为回文结果
	for right := 0; right < strLen; right++ {
		for left := right; left >= 0; left-- {
			if str[right] != str[left] {
				continue
			}

			// 1. 指向同一个元素
			// 2. left和right对应的元素相等 且中间只差一个元素
			// 3. left和right对应的元素相等 且只看left与right区间内元素时是回文的
			// 以上三种情况认为该子串是回文的
			if right == left || (right-left == 1) || dp[right-1][left+1] {
				dp[right][left] = true
			}
		}
	}
}
