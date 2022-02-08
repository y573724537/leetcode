package main

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

const pick = 1

func guessNumber(n int) int {
	low, high := 1, n

	for low <= high {
		mid := low + (high-low)>>1
		result := guess(mid)
		if result == 0 {
			return mid
		} else if result == -1 {
			high = mid - 1
		} else if result == 1 {
			low = mid + 1
		}
	}

	return 1
}

func guess(num int) int {
	if num == pick {
		return 0
	}

	if num < pick {
		return 1
	}

	return -1
}
