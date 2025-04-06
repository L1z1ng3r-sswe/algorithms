func canJump(nums []int) bool {
	target := len(nums) - 1

	dp := make(map[int]struct{}, len(nums))

	var dfs func(idx int) bool
	dfs = func(idx int) bool {
		if idx >= target {
			return true
		}

		for i := nums[idx]; i > 0; i-- {
			if _, ok := dp[idx+i]; !ok && dfs(idx+i) {
				return true
			}
		}

		dp[idx] = struct{}{}
		return false
	}

	return dfs(0)
}

// time: O(n^2)
// space: O(n)

// greedy

func canJump(nums []int) bool {
	n := len(nums)
	target := n - 1

	for i := n - 2; i >= 0; i-- {
		num := nums[i]

		if i+num >= target {
			target = i
		}
	}

	return target == 0
}

// time: O(n)