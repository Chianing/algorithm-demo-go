package tree

// 和为k的子数组个数
// https://leetcode.cn/problems/subarray-sum-equals-k/
func subarraySum(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}

	ret := 0
	preSum := 0
	preSumMap := map[int]int{0: 1}

	for _, v := range nums {
		preSum += v
		if preSumMap[preSum-k] > 0 {
			ret += preSumMap[preSum-k]
		}
		preSumMap[preSum]++
	}

	return ret

}
