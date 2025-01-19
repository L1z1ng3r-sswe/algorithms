func longestConsecutive(nums []int) int {
	numSet := make(map[int]struct{})
	for _, num := range nums {
		numSet[num] = struct{}{}
	}

	var maxStreak int

	for num, _ := range numSet {
		if _, ok := numSet[num-1]; !ok { // must be start of the sequence
			var streak int

			for {
				if _, ok := numSet[num]; !ok {
					break
				}

				num++
				streak++
			}

			maxStreak = max(maxStreak, streak)
		}
	}

	return maxStreak
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// n - distinct nums
// N - len(nums)

// time: O(N)
// space: O(n)