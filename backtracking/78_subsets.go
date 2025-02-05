func subsets(nums []int) [][]int {
	res := make([][]int, 0, 2^len(nums))
	dfs(nums, 0, []int{}, &res)
	return res
}

func dfs(nums []int, idx int, currSet []int, res *[][]int) {
	if idx == len(nums) {
		*res = append(*res, append([]int{}, currSet...))
		return
	}

	dfs(nums, idx+1, currSet, res)
	dfs(nums, idx+1, append(currSet, nums[idx]), res)
}

// time: 2^n