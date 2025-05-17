func canPartitionKSubsets(nums []int, k int) bool {
	var totalSum int
	for _, num := range nums {
		totalSum += num
	}
	if totalSum%k != 0 {
		return false
	}
	target := totalSum / k

	var backtrack func(currSum int, idx int, k int) bool
	backtrack = func(currSum int, idx int, k int) bool {
		if k == 1 {
			return true
		}

		if currSum == target {
			return backtrack(0, 0, k-1)
		}

		if currSum > target || idx == len(nums) {
			return false
		}

		if nums[idx] == 0 {
			return backtrack(currSum, idx+1, k)
		}

		temp := nums[idx]

		nums[idx] = 0

		if backtrack(currSum+temp, idx+1, k) {
			return true
		}

		nums[idx] = temp

		return backtrack(currSum, idx+1, k)
	}

	return backtrack(0, 0, k)
}