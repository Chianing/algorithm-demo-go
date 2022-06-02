package permutation

import "sort"

// 组合求和
// https://leetcode.cn/problems/combination-sum-ii/
func combinationSum2(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}

	ret := make([][]int, 0)
	sort.Ints(candidates)

	var backtrace func(arr, subRet []int, index, target int, ret *[][]int)
	backtrace = func(arr, subRet []int, index, target int, ret *[][]int) {
		if target < 0 {
			return
		}
		if target == 0 {
			subRetCopy := make([]int, len(subRet))
			copy(subRetCopy, subRet)
			*ret = append(*ret, subRetCopy)
			return
		}

		for i := index; i < len(arr); i++ {
			if (i > index) && arr[i] == arr[i-1] {
				// 避免重复
				continue
			}
			subRet = append(subRet, arr[i])
			backtrace(candidates, subRet, i+1, target-arr[i], ret)
			subRet = subRet[:len(subRet)-1]
		}

	}

	backtrace(candidates, make([]int, 0), 0, target, &ret)

	return ret
}
