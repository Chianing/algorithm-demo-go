package other

import "demo/util"

// 找到数组中消失的所有数字
// https://leetcode.cn/problems/find-all-numbers-disappeared-in-an-array/
func findDisappearedNumbers(nums []int) []int {
	ret := make([]int, 0)
	if len(nums) == 0 {
		return ret
	}

	for i := 0; i < len(nums); {
		// 当前元素出现在它该出现的位置，无需交换
		if nums[i] == i+1 {
			i++
			continue
		}
		// idealIdx：当前元素应该出现的位置
		idealIdx := nums[i] - 1
		// 当前元素=它理应出现的位置上的现有元素，说明重复了
		if nums[i] == nums[idealIdx] {
			i++
			continue
		}

		// 不重复，进行交换
		util.Swap(nums, i, idealIdx)
	}

	for i, v := range nums {
		if v != i+1 {
			ret = append(ret, i+1)
		}
	}

	return ret

}
