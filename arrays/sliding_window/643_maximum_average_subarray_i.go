func findMaxAverage(nums []int, k int) float64 {
	var maxVal int
	for _, num := range nums[:k] {
		maxVal += num
	}

	currVal := maxVal
	for i := k; i < len(nums); i++ {
		currVal -= nums[i-k]
		currVal += nums[i]
		maxVal = max(maxVal, currVal)
	}

	return float64(maxVal) / float64(k)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}