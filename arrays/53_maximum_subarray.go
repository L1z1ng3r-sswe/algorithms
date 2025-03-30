func maxSubArray(nums []int) int {
	res := math.MinInt64

	var currSum int
	for _, num := range nums {
		currSum = max(currSum+num, num)
		res = max(res, currSum)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}