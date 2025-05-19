func canPartitionKSubsets(nums []int, k int) bool {
	var totalSum int
	for _, num := range nums {
		totalSum += num
	}
	if totalSum%k != 0 {
		return false
	}
	target := totalSum / k

	sort.Ints(nums)

	buckets := make([]int, k)

	var backtrack func(idx int) bool
	backtrack = func(idx int) bool {
		if idx == -1 {
			for i := range k {
				if buckets[i] != target {
					return false
				}
			}

			return true
		}

		for i := range k {
			if buckets[i]+nums[idx] > target {
				continue
			}

			buckets[i] += nums[idx]

			if backtrack(idx - 1) {
				return true
			}

			buckets[i] -= nums[idx] // backtrack

			if buckets[i] == 0 {
				break
			}
		}

		return false
	}

	return backtrack(len(nums) - 1)
}

// time: O(k ^ n)
// space (k + n) k for buckets, n for max recursion depth

func canPartitionKSubsets(nums []int, k int) bool {
	var totalSum int
	for _, num := range nums {
		totalSum += num
	}
	if totalSum%k != 0 {
		return false
	}
	target := totalSum / k

	sort.Sort(sort.Reverse(sort.IntSlice(nums))) // Descending order optimization
	used := make([]bool, len(nums))

	var backtrack func(idx int, currSum int, kLeft int) bool
	backtrack = func(idx int, currSum int, kLeft int) bool {
		// basecase
		if kLeft == 1 {
			return true
		}

		// if the current group is filled - start filling the next group
		if currSum == target {
			return backtrack(0, 0, kLeft-1)
		}

		if idx >= len(nums) || currSum > target {
			return false
		}

		for i := idx; i < len(nums); i++ {
			if !used[i] {
				used[i] = true
				if backtrack(i+1, currSum+nums[i], kLeft) {
					return true
				}
				used[i] = false

				// optimization: the current number can not be placed anywhere - stop recursion
				if currSum == 0 {
					break
				}

				// optimization: skip duplicate numbers to avoid redundant work
				for i < len(nums)-1 && nums[i] == nums[i+1] {
					i++
				}
			}
		}

		return false
	}

	return backtrack(0, 0, k)
}

// time: O(k(2^n))
// space: O(n)

