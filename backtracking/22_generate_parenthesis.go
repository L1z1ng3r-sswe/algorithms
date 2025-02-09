func generateParenthesis(n int) []string {
	var res []string

	var backtrack func(s []byte, open, close int)
	backtrack = func(s []byte, open, close int) {
		if open == 0 && close == 0 {
			res = append(res, string(s))
			return
		}

		if open > 0 {
			backtrack(append(s, '('), open-1, close+1)
		}

		if close > 0 {
			backtrack(append(s, ')'), open, close-1)
		}
	}
	backtrack([]byte{}, n, 0)

	return res
}