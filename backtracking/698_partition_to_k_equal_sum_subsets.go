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
