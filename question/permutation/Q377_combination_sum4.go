package permutation

// 组合总和
// https://leetcode.cn/problems/combination-sum-iv/
func combinationSum4(nums []int, target int) int {
	length := len(nums)
	if length <= 0 || target < 0 {
		return 0
	}

	// dp[i]代表 target为多少时 有多少种组合
	dp := make([]int, target+1)
	dp[0] = 1

	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if i >= num {
				dp[i] += dp[i-num]
			}
		}
	}

	return dp[target]
}
