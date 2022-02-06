package main

import "fmt"

func searchRange(nums []int, target int) []int {
	return []int{leftBound(nums, target), rightBound(nums, target)}
}

func leftBound(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	low, high := 0, len(nums)-1

	for low <= high {
		mid := low + (high-low)>>1

		if nums[mid] < target {
			low = mid + 1
		} else if target < nums[mid] {
			high = mid - 1
		} else {
			high = mid - 1
		}
	}

	if low == len(nums) || nums[low] != target {
		return -1
	}

	return low
}

func rightBound(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	low, high := 0, len(nums)-1

	for low <= high {
		mid := low + (high-low)>>1

		if target < nums[mid] {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			low = mid + 1
		}
	}

	if high < 0 || nums[high] != target {
		return -1
	}

	return high
}

func main() {
	fmt.Printf("%+v\n", searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
}
