package main

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	low, high := 1, n

	for low <= high {
		mid := low + (high-low)>>1

		if !isBadVersion(mid) {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return low
}

func isBadVersion(version int) bool {
	return true
}
