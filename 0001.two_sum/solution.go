package solution

func twoSum(nums []int, target int) []int {
	dict := map[int]int{}

	for i := 0; i < len(nums); i++ {
		anothor := target - nums[i]
		if j, ok := dict[anothor]; ok {
			return []int{j, i}
		}

		dict[nums[i]] = i
	}

	return nil
}
