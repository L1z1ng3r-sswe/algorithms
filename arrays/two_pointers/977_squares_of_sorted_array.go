func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))
	left, right := 0, len(nums)-1

	idx := len(nums) - 1
	for left <= right {
		sqLeft, sqRight := nums[left]*nums[left], nums[right]*nums[right]

		if sqLeft > sqRight {
			res[idx] = sqLeft
			left++
		} else {
			res[idx] = sqRight
			right--
		}

		idx--
	}

	return res
}