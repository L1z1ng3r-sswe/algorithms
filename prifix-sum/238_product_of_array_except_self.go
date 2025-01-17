func productExceptSelf(nums []int) []int {
	n := len(nums)

	pref := make([]int, n+1)
	pref[0] = 1

	for i, _ := range nums {
		pref[i+1] = pref[i] * nums[i]
	}

	res := make([]int, len(nums))

	suff := 1
	for i := n - 1; i >= 0; i-- {
		res[i] = pref[i] * suff
		suff *= nums[i]
	}

	return res
}