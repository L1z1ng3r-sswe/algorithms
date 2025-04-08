func findNumberOfLIS(nums []int) int {
	n := len(nums)

	dp := make([]int, n)
	opts := make([]int, n)
	for i := range dp {
		dp[i] = 1
		opts[i] = 1
	}

	lis := 1
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				if dp[i] < dp[j]+1 {
					dp[i] = dp[j] + 1
					opts[i] = opts[j]
				} else if dp[j]+1 == dp[i] { // another option
					opts[i] += opts[j]
				}
			}
		}

		if lis < dp[i] {
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

// time: O(n^2)
// space: O(n)