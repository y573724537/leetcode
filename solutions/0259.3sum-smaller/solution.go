package solution

import "sort"

func threeSumSmaller(nums []int, target int) int {
	sort.Ints(nums)

	triNum := 0

	for first := 0; first < len(nums); first++ {
		second := first + 1
		third := len(nums) - 1

		for second < third {
			sum := nums[first] + nums[second] + nums[third]
			if sum < target {
				triNum += third - second
				second++
			} else {
				third--
			}
		}
	}

	return triNum
}
