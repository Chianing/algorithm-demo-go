package strings

// 最长回文子串
// https://leetcode.cn/problems/longest-palindromic-substring/
func longestPalindrome(s string) string {
	length := len(s)
	if length <= 1 {
		return s
	}

	ret := s[:1]
	retLen := 1

	dp := make([][]bool, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]bool, length)
	}

	for right := 0; right < length; right++ {
		for left := right; left >= 0; left-- {
			if left == right {
				dp[left][right] = true
				continue
			}
			if s[left] != s[right] {
				continue
			}
			if (right-left == 1) || dp[left+1][right-1] {
				dp[left][right] = true
				newLen := right - left + 1
				if newLen > retLen {
					ret = s[left : right+1]
					retLen = newLen
				}
			}
		}

	}

	return ret

}
