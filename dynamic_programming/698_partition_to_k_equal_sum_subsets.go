func canPartitionKSubsets(nums []int, k int) bool {
	target, ok := getTarget(nums, k)
	if !ok {
		return false
	}

	buckets := make([]int, k)

	var backtrack func(idx int) bool
	backtrack = func(idx int) bool {
		if idx == len(nums) {
			for i := 0; i < k; i++ {
				if buckets[i] != target {
					return false
				}

			}
			return true
		}

		num := nums[idx]
		for i := 0; i < k; i++ {
			if buckets[i]+num <= target {
				buckets[i] += num
				if backtrack(idx + 1) {
					return true
				}
				buckets[i] -= num

				if buckets[i] == 0 {
					break
				}
			}

		}

		return false
	}

	return backtrack(0)
}

func getTarget(nums []int, k int) (int, bool) {
	var totalSum int
	for _, num := range nums {
		totalSum += num
	}

	return totalSum / k, totalSum%k == 0
}

// time: O(K^N)
// space: O(K)