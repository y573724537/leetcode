package main

func mySqrt(x int) int {
	low, high := 0, x

	for low <= high {
		mid := low + (high-low)>>1
		pow := mid * mid

		if pow == x {
			return mid
		} else if x < pow {
			high = mid - 1
		} else if pow < x {
			low = mid + 1
		}
	}

	return high
}
