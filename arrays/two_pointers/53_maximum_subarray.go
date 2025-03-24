func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	start, end := 0, len(nums)-1

	var res int

	for _, num := range nums {
		res += num
	}

	for start < end {
		if nums[start] < nums[end] {
			res = max(res, res-nums[start])
			start++
		} else {
			res = max(res, res-nums[end])
			end--
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}