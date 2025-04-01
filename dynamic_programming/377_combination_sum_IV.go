func combinationSum4(nums []int, target int) int {
	dp := make(map[int]int)

	var dfs func(sum int) int
	dfs = func(sum int) int {
		if sum == target {
			return 1
		}

		if val, ok := dp[sum]; ok {
			return val
		}

		res := 0

		for _, num := range nums {
			if sum+num <= target {
				res += dfs(sum + num)
			}
		}

		dp[sum] = res
		return res
	}

	return dfs(0)
}

// T - target
// N - len(nums)

// time: O(T*N)
// space: O(T)