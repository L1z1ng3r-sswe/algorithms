func findDuplicates(nums []int) []int {
	var res []int

	for _, num := range nums {
		idx := abs(num) - 1
		nums[idx] = -nums[idx]
		if nums[idx] > 0 {
			res = append(res, idx+1)
		}
	}

	return res
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}