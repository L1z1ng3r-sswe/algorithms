type key struct {
	isToBuy bool
	profit  int
	idx     int
}

func maxProfit(prices []int) int {
	3
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

func maxProfit(prices []int) int {
	dp := make(map[[2]int]int) // cache

	var dfs func(i int, canBuy int) int
	dfs = func(i int, canBuy int) int {
		if i >= len(prices) {
			return 0
		}

		pair := [2]int{i, canBuy}
		if v, ok := dp[pair]; ok {
			return v
		}

		var res int
		if canBuy == 1 {
			wait := dfs(i+1, 1)
			buy := -prices[i] + dfs(i+1, 0)
			res = max(wait, buy)
		} else {
			wait := dfs(i+1, 0)
			sell := prices[i] + dfs(i+2, 1)
			res = max(wait, sell)
		}

		dp[pair] = res
		return res
	}

	return dfs(0, 1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// time: O(n)
// space: O(n)

func maxProfit(prices []int) int {
	n := len(prices)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = []int{-1, -1}
	}

	var dp func(day int, holding int) int
	dp = func(day int, holding int) int {
		if day >= n {
			return 0
		}
		if memo[day][holding] != -1 {
			return memo[day][holding]
		}

		if holding == 1 {
			// Option 1: Sell today -> cooldown tomorrow
			sell := prices[day] + dp(day+2, 0)
			// Option 2: Hold
			hold := dp(day+1, 1)
			memo[day][holding] = max(sell, hold)
		} else {
			// Option 1: Buy today
			buy := -prices[day] + dp(day+1, 1)
			// Option 2: Skip
			skip := dp(day+1, 0)
			memo[day][holding] = max(buy, skip)
		}
		return memo[day][holding]
	}

	return dp(0, 0)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}