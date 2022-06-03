package strings

// Z字型变换
// https://leetcode.cn/problems/zigzag-conversion/
func convert(s string, numRows int) string {
	if len(s) == 0 || numRows < 2 {
		return s
	}

	convertedArr := make([]string, numRows)

	convertIndex := 0
	flag := -1

	for _, v := range s {
		convertedArr[convertIndex] += string(v)
		if convertIndex == 0 || convertIndex == numRows-1 {
			flag = -flag
		}
		convertIndex += flag
	}

	ret := ""
	for _, v := range convertedArr {
		ret += v
	}

	return ret
}
