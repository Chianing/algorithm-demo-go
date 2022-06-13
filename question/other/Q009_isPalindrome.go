package other

// 回文数
// https://leetcode.cn/problems/palindrome-number/
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x == 0 || x < 10 {
		return true
	}
	if x%10 == 0 {
		return false
	}

	subNum := 0
	for x > subNum {
		subNum *= 10
		subNum += x % 10
		x = x / 10
	}

	return subNum == x || subNum/10 == x

}
