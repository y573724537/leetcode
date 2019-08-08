package solution

func mySqrt(x int) int {
	if x == 0 {
		return x
	}

	var mid int
	low, high := 1, x
	for low <= high {
		mid = (low + high) / 2

		if mid == x/mid {
			return mid
		}

		if mid > x/mid {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return high
}
