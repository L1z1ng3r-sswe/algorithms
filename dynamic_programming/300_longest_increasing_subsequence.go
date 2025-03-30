func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	n := len(nums)
	dp := make([]int, n)

	for i := range dp {
		dp[i] = 1
	}

	maxLength := 1

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		maxLength = max(maxLength, dp[i])
	}

	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// binary approach
// ...