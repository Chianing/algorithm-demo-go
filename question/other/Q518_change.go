package other

// 零钱兑换
// https://leetcode.cn/problems/coin-change-2/
func change(amount int, coins []int) int {
	dp := make([]int, amount+1)

	dp[0] = 1
	for _, v := range coins {
		for j := 1; j <= amount; j++ {
			if j < v {
				continue
			}
			dp[j] += dp[j-v]
		}
	}

	return dp[amount]

}
