package other

import "demo/util"

// 缺失的第一个正数
// https://leetcode.cn/problems/first-missing-positive/
func firstMissingPositive(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 1
	}

	for i := 0; i < length; i++ {
		// 1. 当前元素需要是正数
		// 2. 当前元素需要小于数组长度
		// 3. 当前元素若看作索引的话 替换的前提是该位置上现有元素不该在那个位置
		// 如[1, 1, 3]数组中 第二个1位置
		for nums[i] > 0 && nums[i] < length && nums[nums[i]-1] != nums[i] {
			util.Swap(nums, nums[i]-1, i)
		}
	}

	for i := 0; i < length; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	return length + 1

}
