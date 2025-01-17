func isSubsequence(s string, t string) bool {
	var m, n = len(s), len(t)
	if m == 0 {
		return true
	}

	var currIdx int // 0
	for i := 0; i < n; i++ {
		if t[i] == s[currIdx] {
			currIdx++
		}
		if currIdx == m {
			return true
		}
	}
	return false
}

