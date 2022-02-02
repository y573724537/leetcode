package solution

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	low, high := 0, len(nums)-1

	for low <= high {
		mid := (low + high) / 2

		if target < nums[mid] {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			return mid
		}
	}

	return low
}
