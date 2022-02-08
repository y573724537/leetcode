package main

func findPeakElement(nums []int) int {
	low, high := 0, len(nums)-1

	for low <= high {
		mid := low + (high-low)>>1

		if mid == 0 {
			if high == 0 || nums[mid+1] < nums[mid] {
				return mid
			} else {
				low = mid + 1
			}
		} else if mid == len(nums)-1 {
			if low == len(nums)-1 || nums[mid-1] < nums[mid] {
				return mid
			} else {
				high = mid - 1
			}
		} else {
			if nums[mid-1] < nums[mid] && nums[mid+1] < nums[mid] {
				return mid
			} else if nums[mid] < nums[mid-1] {
				high = mid - 1
			} else if nums[mid] < nums[mid+1] {
				low = mid + 1
			}
		}
	}

	return -1
}
