func permute(nums []int) [][]int {
	res := [][]int{}

	var dfs func(start int) 
	dfs = func(start int) {
		if start == len(nums) {
			res = append(res, append([]int{}, nums...))
			return
		}

		for i := start; i < len(nums); i++ {
			nums[i], nums[start] = nums[start], nums[i]
			dfs(start+1)
			nums[i], nums[start] = nums[start], nums[i]
		}
	}
	
	dsf(0)
	
	return res
}


[1,2,3] -> 1,2,3  1,3,2     2,1,3    2,3,1    3,1,2    3,2,1


time complexity: O(N!)