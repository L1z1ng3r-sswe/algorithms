func indexPairsOfString(str string, words []string) [][]int {
	var res [][]int
	n := len(str)
	for _, word := range words {
		m := len(word)
		if m <= n {
			for i := 0; i+m <= n; i++ {
				if word == str[i:i+m] {
					res = append(res, []int{i, i + m - 1})
				}
			}
		}
	}

	return res
}
