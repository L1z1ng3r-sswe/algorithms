func longestPalindrome(s string) string {
	n := len(s)
	if n == 1 {
		return s
	}

	res := 1
	resS := s[:1]
	for i := 0; i < n-1; i++ {
		// odd
		l, r := i, i

		for l >= 0 && r < n && s[l] == s[r] {
			if res < r-l+1 {
				res = r - l + 1
				resS = s[l : r+1]
			}

			l--
			r++
		}

		l, r = i, i+1
		for l >= 0 && r < n && s[l] == s[r] {
			if res < r-l+1 {
				res = r - l + 1
				resS = s[l : r+1]
			}

			l--
			r++
		}
	}

	return resS
}