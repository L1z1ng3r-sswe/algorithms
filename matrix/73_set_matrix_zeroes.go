func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	var isFirstRowZero bool // false

	for c := 0; c < n; c++ {
		if matrix[0][c] == 0 {
			isFirstRowZero = true
			break
		}
	}

	for r := 1; r < m; r++ {
		for c := 0; c < n; c++ {
			if matrix[r][c] == 0 {
				matrix[0][c] = 0
				matrix[r][0] = 0
			}
		}
	}

	for r := 1; r < m; r++ {
		if matrix[r][0] == 0 {
			for c := 1; c < n; c++ {
				matrix[r][c] = 0
			}
		}
	}

	for c := 0; c < n; c++ {
		if matrix[0][c] == 0 {
			for r := 1; r < m; r++ {
				matrix[r][c] = 0
			}
		}
	}

	if isFirstRowZero {
		for c := 0; c < n; c++ {
			matrix[0][c] = 0
		}
	}
}