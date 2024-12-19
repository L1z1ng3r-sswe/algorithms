func nextGreatestLetter(letters []byte, target byte) byte {
	left, right := 0, len(letters)-1

	for left <= right {
		mid := left + (right-left)/2

		if letters[mid] == target || letters[mid] < target {
			left = mid + 1
		} else if letters[mid] > target && mid > 0 && letters[mid-1] > target {
			right = mid - 1
		} else {
			return letters[mid]
		}
	}

	return letters[0]
}