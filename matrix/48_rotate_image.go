func rotate(mx [][]int) {
	left, right := 0, len(mx)-1

	for left < right { // n / 2 (i.e. for 3 / 2 = 1, 4 / 2 = 2)
		for i := 0; i < right-left; i++ {
			top, bottom := left, right

			leftTp := mx[top][left+i]

			mx[top][left+i] = mx[bottom-i][left]
			mx[bottom-i][left] = mx[bottom][right-i]
			mx[bottom][right-i] = mx[top+i][right]
			mx[top+i][right] = leftTp
		}

		left++
		right--
	}
}