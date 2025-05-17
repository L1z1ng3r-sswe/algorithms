/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	left, right := 1, n

	for left < right {
		mid := left + (right-left)/2

		res := guess(mid)
		switch res {
		case -1: // too large
			right = mid - 1
		case 1: // too little
			left = mid + 1
		default:
			return mid
		}
	}

	return left
}

// time: O(log(n))