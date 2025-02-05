func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)

	resSet := make(map[string][]int)
	dfs(nums, []byte{}, []int{}, 0, resSet)

	var res [][]int
	for _, subset := range resSet {
		res = append(res, subset)
	}

	return res
}

func dfs(nums []int, subsetBytes []byte, subset []int, idx int, resSet map[string][]int) {
	if idx == len(nums) {
		resSet[string(subsetBytes)] = append([]int{}, subset...)
		return
	}

	dfs(nums, subsetBytes, subset, idx+1, resSet)
	dfs(nums, append(subsetBytes, '-', byte(nums[idx])), append(subset, nums[idx]), idx+1, resSet)
}