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
	used := make([]bool, len(nums))
	sort.Sort(sort.Reverse(sort.IntSlice(nums))) // Descending order optimization

	var backtrack func(idx int, currSum int, k int) bool
	backtrack = func(idx int, currSum int, k int) bool {
		if k == 1 {
			return true
		}

		if currSum == target {
			return backtrack(0, 0, k-1)
		}

		for i := idx; i < len(nums); i++ {
			if !used[i] && currSum+nums[i] <= target {
				used[i] = true
				if backtrack(i+1, currSum+nums[i], k) {
					return true
				}
				used[i] = false

				if currSum == 0 {
					break
				}

				for i+1 < len(nums) && nums[i] == nums[i+1] {
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

