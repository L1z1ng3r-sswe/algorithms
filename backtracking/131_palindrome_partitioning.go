
func partition(s string) [][]string {
	var n = len(s)

	var isPalindrome = func(start int, end int) bool {
		for i, j := start, end-1; i < j; i, j = i+1, j-1 {
			if s[i] != s[j] {
				return false
			}
		}
		return true
	}

	var res [][]string

	var backtrack func(start int, substrings []string)
	backtrack = func(start int, substrings []string) {
		if start >= n {
			res = append(res, append([]string{}, substrings...))
			return
		}

		for end := start + 1; end <= n; end++ {
			if isPalindrome(start, end) {
				backtrack(end, append(substrings, s[start:end]))
			}
		}
	}
	backtrack(0, []string{})

	return res
}
