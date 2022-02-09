package main

func nextGreatestLetter(letters []byte, target byte) byte {
	left := leftBound(letters, target)
}

func leftBound(letters []byte, target byte) int {
	low, high := 0, len(letters)-1

	for low <= high {
		mid := low + (high-low)>>1

		if letters[mid] == target {
			high = mid - 1
		}

		if letters[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return low
}

func isSmaller(a, b byte) bool {

}
