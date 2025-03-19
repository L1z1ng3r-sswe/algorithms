func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	maxVal := max(nums[0], nums[1])

	if len(nums) == 2 {
		return maxVal
	}

	prevMaxVal := nums[0]

	for i := 2; i < len(nums); i++ {
		newMaxVal := max(prevMaxVal+nums[i], maxVal)
		prevMaxVal = maxVal
		maxVal = newMaxVal
	}

	return maxVal
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// time: O(n)
// space: O(n)