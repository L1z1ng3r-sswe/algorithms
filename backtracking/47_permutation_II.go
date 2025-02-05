func permuteUnique(nums []int) [][]int {
	res := [][]int{}
	var backtrack func(start int)
	backtrack = func(start int) {
		if start == len(nums) {
			res = append(res, append([]int, nums...))
			return
		}

		visited := make(map[int]struct{})

		for i := start; i < len(nums); i++ {
			if _, ok := visited[nums[i]]; !ok {
				visited[nums[i]] = struct{}{}
				
				nums[i], nums[start] = nums[start], nums[i]
				res = append(start++)
				nums[i], nums[start] = nums[start], nums[i]
			}
		}
	}

	return res
}