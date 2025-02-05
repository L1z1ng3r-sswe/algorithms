func combinationSum3(k int, n int) [][]int {
	var res [][]int

	var backtrack func(start int, target int, comb []int)
	backtrack = func(start int, target int, comb []int) {
		if target == 0 && len(comb) == k {
			res = append(res, append([]int{}, comb...))
			return
		}

		if len(comb) >= k {
			return
		}

		for i := start; i <= 9; i++ {
			if target >= i {
				backtrack(i+1, target-i, append(comb, i))
			}
		}
	}

	backtrack(1, n, []int{})
	return res
}