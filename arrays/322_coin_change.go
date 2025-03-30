func coinChange(coins []int, amount int) int {
	sort.Ints(coins)

	lim := amount + 1
	dp := make([]int, lim)
	for i := range dp {
		dp[i] = lim
	}
	dp[0] = 0

	for i := coins[0]; i <= amount; i++ {
		for _, coin := range coins {
			if i < coin {
				break
			}

			if dp[i-coin] != lim && dp[i-coin]+1 < dp[i] {
				dp[i] = dp[i-coin] + 1
			}
		}
	}

	if dp[amount] == lim {
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

// m - amount
// n - len(coins)
// time: O(m*n)
// space: O(m)