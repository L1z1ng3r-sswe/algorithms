func combine(n int, k int) [][]int {
	var res [][]int

	var backtrack func(start int, comb []int)
	backtrack = func(start int, comb []int) {
		if len(comb) == k {
			res = append(res, append([]int{}, comb...))
			return
		}

		for i := start; i <= n; i++ {
			backtrack(i+1, append(comb, i))
		}
	}

	backtrack(1, []int{})

	return res
}