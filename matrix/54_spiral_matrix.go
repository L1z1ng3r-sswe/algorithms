func spiralOrder(mx [][]int) []int {
	m, n := len(mx), len(mx[0])

	res := make([]int, 0, m*n)
	left, right, top, bottom := 0, n-1, 0, m-1
	for top <= bottom && left <= right { // TODO: determine the condition
		for i := left; i <= right; i++ {
			res = append(res, mx[top][i])
		}
		top++

		for i := top; i <= bottom; i++ {
			res = append(res, mx[i][right])
		}
		right--

		if top <= bottom {
			for i := right; i >= left; i-- {
				res = append(res, mx[bottom][i])
			}
			bottom--
		}

		if right >= left {
			for i := bottom; i >= top; i-- {
				res = append(res, mx[i][left])
			}
			left++
		}
	}

	return res
}