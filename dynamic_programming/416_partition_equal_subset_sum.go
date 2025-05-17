func canPartition(nums []int) bool {
	var totalSum int
	for _, num := range nums {
		totalSum += num
	}
	if totalSum%2 != 0 {
		return false
	}
	target := totalSum / 2

	sort.Ints(nums)
	n := len(nums) - 1

	deadend := make(map[[2]int]struct{})

	var dfs func(idx int, sum int) bool
	dfs = func(idx int, sum int) bool {
		if sum == target {
			return true
		}

		if idx >= n || sum > target {
			return false
		}

		pair := [2]int{idx, sum}
		if _, ok := deadend[pair]; ok {
			return false
		}

		if dfs(idx+1, sum+nums[idx]) {
			return true
		}

		if dfs(idx+1, sum) {
			return true
		}

		deadend[pair] = struct{}{}
		return false
	}

	return dfs(0, 0)
}

// time: O(n*target)
// space: O(n*target)

func canPartition(nums []int) bool { // 1, 5, 11, 5; target = 11
	var totalSum int
	for _, num := range nums {
		totalSum += num
	}
	if totalSum%2 != 0 {
		return false
	}
	target := totalSum / 2

	sort.Ints(nums)

	possibleSums := make([]bool, target+1)
	possibleSums[0] = true

	for _, num := range nums {
		for sum := len(possibleSums) - 1; sum >= 0; sum-- {
			if possibleSums[sum] {
				newSum := sum + num
				if newSum > target {
					continue
				}

				possibleSums[newSum] = true
			}
		}
		if possibleSums[target] {
			return true
		}
	}

	return false
}

// time: O(n * target)
// space: O(target)