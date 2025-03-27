func maxProduct(nums []int) int {
	res := nums[0]
	maxCurr, minCurr := nums[0], nums[0]

	for _, num := range nums[1:] {
		if num < 0 {
			maxCurr, minCurr = minCurr, maxCurr
		}

		maxCurr = max(num, maxCurr*num)
		minCurr = min(num, minCurr*num)

		res = max(res, maxCurr)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}