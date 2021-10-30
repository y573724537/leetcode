package solution

func twoSum(nums []int, target int) []int {
	records := make(map[int]int, len(nums))

	for i := 0; i < len(nums); i++ {
		mate := target - nums[i]
		if mateIdx, ok := records[mate]; ok {
			return []int{mateIdx, i}
		}

		records[nums[i]] = i
	}

	return nil
}
