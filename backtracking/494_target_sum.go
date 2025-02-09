func findTargetSumWays(nums []int, target int) int {
	var n = len(nums)

	dp := make(map[key]int)

	var backtrack func(start int, sum int) int
	backtrack = func(start int, sum int) int {
		if start == n {
			if sum == target {
				return 1
			}
			return 0
		}

		if val, ok := dp[key{start, sum}]; ok {
			return val
		}

		subRes := backtrack(start+1, sum-nums[start])
		addRes := backtrack(start+1, sum+nums[start])
		res := subRes + addRes

		dp[key{start, sum}] = res

		return res
	}

	return backtrack(0, 0)
}

type key struct {
	idx     int
	currSum int
}