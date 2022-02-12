package main

func findKthLargest(nums []int, k int) int {
	for i := 0; i < k; i++ {
		maxIdx := i
		for j := i + 1; j < len(nums); j++ {
			if nums[maxIdx] < nums[j] {
				maxIdx = j
			}
		}

		if maxIdx != i {
			nums[maxIdx], nums[i] = nums[i], nums[maxIdx]
		}
	}

	return nums[k-1]
}
