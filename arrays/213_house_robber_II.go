func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	return max(robber(nums[1:]), robber(nums[:len(nums)-1]))
}

func robber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	prevMaxVal, maxVal := nums[0], max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		newMaxVal := max(maxVal, prevMaxVal+nums[i])
		maxVal, prevMaxVal = newMaxVal, maxVal
	}

	return maxVal
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}