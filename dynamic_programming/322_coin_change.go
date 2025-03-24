func coinChange(coins []int, amount int) int {
	sort.Ints(coins)

	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = math.MaxInt64
	}
	dp[0] = 0

	for i := coins[0]; i <= amount; i++ {
		for _, coin := range coins {
			if i >= coin {
				if dp[i-coin] != math.MaxInt64 {
					dp[i] = min(dp[i], dp[i-coin]+1)
				}
			} else {
				break
			}
		}
	}

	if dp[amount] == math.MaxInt64 {
		return -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}