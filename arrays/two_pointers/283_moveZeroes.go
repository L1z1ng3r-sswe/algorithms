func moveZeroes(nums []int) {
	zeroI := 0

	for i := 0; i < len(nums); i++ {
		num := nums[i]

		if num != 0 {
			nums[i], nums[zeroI] = nums[zeroI], nums[i]
			zeroI++
		}
	}
}