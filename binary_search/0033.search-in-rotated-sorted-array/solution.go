package main

func search(nums []int, target int) int {
	low, high := 0, len(nums)-1

	for low <= high {
		mid := low + (high-low)>>1

		if nums[mid] == target {
			return mid
		}

		if nums[low] <= nums[mid] {
			// [low, mid] is sorted
			if nums[low] <= target && target < nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			// [mid, high] is sorted
			if nums[mid] < target && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}

	return -1
}
