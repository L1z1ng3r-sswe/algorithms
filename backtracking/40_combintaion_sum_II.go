func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	var res [][]int

	var backtrack func(start int, target int, comb []int)
	backtrack = func(start int, target int, comb []int) {
		if target == 0 {
			res = append(res, append([]int{}, comb...))
			return
		}

		var prev int
		for i := start; i < len(candidates); i++ {
			if candidates[i] != prev && target >= candidates[i] {
				backtrack(i+1, target-candidates[i], append(comb, candidates[i]))
			}

			prev = candidates[i]
		}
	}

	backtrack(0, target, []int{})
	return res
}