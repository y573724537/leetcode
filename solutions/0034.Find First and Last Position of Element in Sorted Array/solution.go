package solution

func searchRange(nums []int, target int) []int {
	left := searchMostStart(nums, target)
	if left == len(nums) || nums[left] != target {
		return []int{-1, -1}
	}
	return []int{left, searchMostStart(nums, target+1) - 1}
}

func searchMostStart(nums []int, target int) int {
	low, high := 0, len(nums)-1

	for low <= high {
		mid := (low + high) / 2

		if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return low
}
