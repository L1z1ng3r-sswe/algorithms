func generateParenthesis(n int) []string {
	var res []string

	var backtrack func(rightRemaining, leftRemaining int, paranthesis []byte)
	backtrack = func(rightRemaining, leftRemaining int, paranthesis []byte) {
		if rightRemaining == 0 && leftRemaining == 0 {
			res = append(res, string(paranthesis))
			return
		}

		if rightRemaining > 0 {
			backtrack(rightRemaining-1, leftRemaining+1, append(paranthesis, '('))
		}

		if leftRemaining > 0 {
			backtrack(rightRemaining, leftRemaining-1, append(paranthesis, ')'))
		}
	}

	backtrack(n, 0, []byte{})

	return res
}