package permutation

import "sort"

// 子集
// https://leetcode.cn/problems/subsets-ii/
func subsetsWithDup(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	ret := make([][]int, 0)
	subRet := make([]int, 0)
	sort.Ints(nums)

	var backtrace func(int)
	backtrace = func(index int) {
		// 不选择该元素
		ret = append(ret, append([]int{}, subRet...))

		for i := index; i < len(nums); i++ {
			if (i > index) && (nums[i-1] == nums[i]) {
				// 避免出现重复元素
				continue
			}

			// 选择该元素
			subRet = append(subRet, nums[i])
			backtrace(i + 1)

			// 撤销选择
			subRet = subRet[:len(subRet)-1]
		}

	}

	backtrace(0)

	return ret

}
