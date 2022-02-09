package main

func moveZeroes(nums []int) {
	for sentinel := len(nums); 0 < sentinel; {
		tmpSentinel := sentinel
		sentinel = 0

		for i := 0; i+1 < tmpSentinel; i++ {
			if nums[i] == 0 {
				nums[i], nums[i+1] = nums[i+1], nums[i]
				sentinel = i + 1
			}
		}
	}
}
