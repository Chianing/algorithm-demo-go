package other

// 零钱兑换
// https://leetcode.cn/problems/coin-change/
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	// 金额为i时 硬币最少的组合方案
	dp := make([]int, amount+1)
	dp[0] = 0
	max := amount + 1

	for i := 1; i <= amount; i++ {
		dp[i] = max
		for _, v := range coins {
			if i >= v {
				dp[i] = getMin(dp[i], dp[i-v]+1)
			}
		}
	}

	if dp[amount] == max {
		return -1
	}

	return dp[amount]

}

func getMin(a, b int) int {
	if a > b {
		return b
	}

	return a
}
