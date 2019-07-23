package solution

func maxSubArray(nums []int) int {
	if len(nums) < 1 {
		return 0
	}

	status := make([]int, len(nums))
	status[0] = nums[0]
	max := nums[0]

	for i := 1; i < len(nums); i++ {
		if status[i-1]+nums[i] > nums[i] {
			status[i] = status[i-1] + nums[i]
		} else {
			status[i] = nums[i]
		}

		if max < status[i] {
			max = status[i]
		}
	}

	return max
}
