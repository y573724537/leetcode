package main

func isMajorityElement(nums []int, target int) bool {
	leftIdx := leftBound(nums, target)
	if leftIdx == -1 {
		return false
	}

	rightIdx := rightBound(nums, target)

	return len(nums)>>1 < rightIdx-leftIdx+1
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
