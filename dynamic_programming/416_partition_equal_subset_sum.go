func canPartition(nums []int) bool {
	var sum int
	for _, num := range nums {
		sum += num
	}
	if len(nums) == 1 || sum%2 != 0 {
		return false
	}

	target := sum / 2

	deadend := make(map[[2]int]struct{})

	var dfs func(idx int, target int) bool
	dfs = func(idx int, target int) bool {
		if target == 0 {
			return true
		}

		if idx == len(nums) || target < 0 {
			return false
		}

		end := [2]int{idx, target}
		if _, ok := deadend[end]; ok {
			return false
		}

		if dfs(idx+1, target-nums[idx]) {
			return true
		}

		if dfs(idx+1, target) {
			return true
		}

		deadend[end] = struct{}{}

		return false
	}

	return dfs(0, target)
}

// target = sum/2
// n = len(nums)

// time: O(n*target)
// space: O(n*target)

func canPartition(nums []int) bool {
	var sum int
	for _, num := range nums {
		sum += num
	}

	if sum%2 != 0 {
		return false
	}

	target := sum / 2

	set := make(map[int]struct{})
	set[0] = struct{}{}

	for _, num := range nums {
		res := make([]int, 0, len(set))
		for sum := range set {
			newSum := sum + num

			if newSum == target {
				return true
			} else if newSum < target {
				res = append(res, newSum)
			}
		}

		for _, num := range res {
			set[num] = struct{}{}
		}
	}

	return false
}

// n - len(nums)
// target = sum(nums) / 2

// time: O(n * target)
// space: O(target)