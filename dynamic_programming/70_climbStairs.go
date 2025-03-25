func climbStairs(n int) int {
	dp := make(map[int]int)

	var dfs func(rem int) int
	dfs = func(rem int) int {
		if rem == 0 {
			return 1
		} else if rem < 0 {
			return 0
		}

		if val, ok := dp[rem]; ok {
			return val
		}

		one, two := dfs(rem-1), dfs(rem-2)
		dp[rem] = one + two

		return one + two
	}

	return dfs(n)
}