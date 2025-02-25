func partition(s string) [][]string {
	var res [][]string

	var backtrack func(start int, subset []string)
	backtrack = func(start int, subset []string) {
		if start >= len(s) {
			res = append(res, append([]string{}, subset...))
			return
		}

		for end := start + 1; end <= len(s); end++ {
			if isPalindrome(s[start:end]) {
				backtrack(end, append(subset, s[start:end]))
			}
		}
	}
	backtrack(0, []string{})

	return res
}

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}