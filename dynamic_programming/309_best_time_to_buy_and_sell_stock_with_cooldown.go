type key struct {
	isToBuy bool
	profit  int
	idx     int
}

func maxProfit(prices []int) int {
	dp := make(map[key]int)

	var dfs func(isToBuy bool, profit int, idx int) int
	dfs = func(isToBuy bool, profit int, idx int) int {
		if idx >= len(prices) {
			return profit
		}

		currKey := key{isToBuy, profit, idx}
		if v, ok := dp[currKey]; ok {
			return v
		}

		if isToBuy {
			dp[currKey] = max(dfs(false, profit-prices[idx], idx+1), dfs(true, profit, idx+1))
		} else {
			dp[currKey] = max(dfs(true, profit+prices[idx], idx+2), dfs(false, profit, idx+1))
		}

		return dp[currKey]
	}

	return dfs(true, 0, 0)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// time: O(2 * n * totalSum)
// space: O(2 * n * totalSum)