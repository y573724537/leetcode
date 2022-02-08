package main

import "fmt"

func findClosestElements(arr []int, k int, x int) []int {
	if k == len(arr) {
		return arr
	}

	position := binarySearch(arr, x)
	if position == 0 {
		return arr[:k]
	}

	if position == len(arr) {
		return arr[len(arr)-k:]
	}

	left, right := 0, len(arr)-1
	if 0 < position-k-1 {
		left = position - k - 1
	}

	if position+k+1 < len(arr)-1 {
		right = position + k + 1
	}

	for k < right-left+1 {
		if arr[right]-x < x-arr[left] {
			left++
		} else {
			right--
		}
	}

	return arr[left : right+1]
}

func binarySearch(arr []int, x int) int {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := low + (high-low)>>1

		if arr[mid] < x {
			low = mid + 1
		} else if x < arr[mid] {
			high = mid - 1
		} else {
			high = mid - 1
		}
	}

	return low
}

func main() {
	fmt.Printf("%v\n", findClosestElements([]int{0, 1, 1, 1, 2, 3, 6, 7, 8, 9}, 9, 4))
}
