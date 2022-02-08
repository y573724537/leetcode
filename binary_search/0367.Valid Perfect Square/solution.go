package main

func isPerfectSquare(num int) bool {
	low, high := 0, num

	for low <= high {
		mid := low + (high-low)>>1
		power := mid * mid

		if power == num {
			return true
		}

		if power < num {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return false
}
