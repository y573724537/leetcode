package main

/**
 * // This is the ArrayReader's API interface.
 * // You should not implement it, or speculate about its implementation
 * type ArrayReader struct {
 * }
 *
 * func (this *ArrayReader) get(index int) int {}
 */

type ArrayReader struct {
}

func (this *ArrayReader) get(index int) int {
	return -1
}

func search(reader ArrayReader, target int) int {
	low, high := 0, 10*10*10*10-1

	for low <= high {
		mid := low + (high-low)>>1

		val := reader.get(mid)

		if target == val {
			return mid
		}

		if target < val {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}
