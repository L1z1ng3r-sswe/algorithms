func countSubstrings(s string) int {
	n := len(s)

	var res int
	for i := 0; i < n; i++ {
		l, r := i, i
		for l >= 0 && r < n && s[l] == s[r] {
			res++
			l--
			r++
		}

		l, r = i, i+1
		for l >= 0 && r < n && s[l] == s[r] {
			res++
			l--
			r++
		}
	}

	return res
}

// time: O(n^2)