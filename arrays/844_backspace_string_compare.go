func backspaceCompare(s string, t string) bool {
	return remaker(s) == remaker(t)
}

func remaker(s string) string {
	var res strings.Builder

	var count int
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != '#' {
			if count > 0 {
				count--
			} else {
				res.WriteByte(s[i])
			}
		} else {
			count++
		}
	}

	return res.String()
}