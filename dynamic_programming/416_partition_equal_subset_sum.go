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

func canPartition(nums []int) bool { // 1, 5, 11, 5; target = 11
	target, ok := getTarget(nums)
	if !ok {
		return false
	}

	sort.Ints(nums)

	possibleSums := make([]bool, target+1)
	possibleSums[0] = true // no nums are needed to sum up 0

	for _, num := range nums {
		for s := target; s >= num; s-- {
			if possibleSums[s-num] {
				possibleSums[s] = true
			}
		}
		if possibleSums[target] {
			return true
		}
	}

	return false
}

func getTarget(nums []int) (int, bool) {
	var totalSum int
	for _, num := range nums {
		totalSum += num
	}
	return totalSum / 2, totalSum%2 == 0
}

// n - len(nums)
// target = sum(nums) / 2

// time: O(n * target)
// space: O(target)