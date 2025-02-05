func findTargetSumWays(nums []int, target int) int {
	dp := make(map[Pair]int)

	var backtrack func(idx int, target int) int
	backtrack = func(idx int, target int) int {
		if idx >= len(nums) {
			if target == 0 {
				return 1
			}
			return 0
		}

		if val, ok := dp[Pair{idx, target}]; ok {
			return val
		}

		pRes := backtrack(idx+1, target-nums[idx])
		mRes := backtrack(idx+1, target+nums[idx])

		dp[Pair{idx + 1, target - nums[idx]}] = pRes
		dp[Pair{idx + 1, target + nums[idx]}] = mRes

		return pRes + mRes
	}

	return backtrack(0, target)
}

type Pair struct {
	Idx    int
	Target int
}