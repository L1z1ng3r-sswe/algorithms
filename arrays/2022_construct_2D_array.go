func construct2DArray(original []int, m, n int) [][]int {
	if len(original) != m*n {
		return [][]int{}
	}

	var res [][]int
	for i := 0; i < len(original); i += n {
		res = append(res, original[i:i+n])
	}
	return res
}