package permutation

// 组合求和
// https://leetcode.cn/problems/combination-sum-iii/
func combinationSum3(k int, n int) [][]int {
	if k <= 0 || k >= 10 || n > 45 || n <= 0 {
		return [][]int{}
	}

	ret := make([][]int, 0)
	subRet := make([]int, 0)

	var backtrace func(int, int)
	backtrace = func(index int, target int) {
		if target < 0 {
			return
		}
		if len(subRet) == k {
			if target == 0 {
				ret = append(ret, append([]int{}, subRet...))
			}
			return
		}

		for i := index; i < 10; i++ {
			subRet = append(subRet, i)
			backtrace(i+1, target-i)
			subRet = subRet[:len(subRet)-1]
		}
	}
	backtrace(1, n)

	return ret

}
