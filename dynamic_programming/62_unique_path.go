func uniquePaths(m int, n int) int {
	row := make([]int, n)
	col := make([]int, m)
	for i := range n {
		row[i] = 1
	}
	for i := range m {
		col[i] = 1
	}

	for r := m - 2; r >= 0; r-- {
		for c := n - 2; c >= 0; c-- {
			cell := row[c] + col[r]
			row[c], col[r] = cell, cell
		}
	}

	return col[0]
}

// time: O(m*n)
// space: O(m+n)