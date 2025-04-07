func findNumberOfLIS(nums []int) int {
	n := len(nums)
	dp, opts := make([]int, n), make([]int, n)
	for i := range opts {
		dp[i], opts[i] = 1, 1
	}

	lis := 1
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				if dp[i] < dp[j]+1 {
					dp[i] = dp[j] + 1
					opts[i] = opts[j]
				} else if dp[i]-1 == dp[j] {
					opts[i] += opts[j]
				}
			}
		}
		if dp[i] > lis {
			lis = dp[i]
		}
	}

	var res int
	for i, l := range dp {
		if l == lis {
			res += opts[i]
		}
	}

	return res
}