package solution

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)

	result := make([][]int, 0)

	for first := 0; first < len(nums); first++ {
		if 0 < first && nums[first-1] == nums[first] {
			continue
		}

		for second := first + 1; second < len(nums); second++ {
			if first+1 < second && nums[second-1] == nums[second] {
				continue
			}

			third := second + 1
			fourth := len(nums) - 1

			for third < fourth {
				if second+1 < third && nums[third-1] == nums[third] {
					third++
					continue
				}

				if fourth < len(nums)-1 && nums[fourth] == nums[fourth+1] {
					fourth--
					continue
				}

				sum := nums[first] + nums[second] + nums[third] + nums[fourth]

				if sum == target {
					result = append(result, []int{nums[first], nums[second], nums[third], nums[fourth]})
					third++
					fourth--
					continue
				}

				if sum < target {
					third++
					continue
				}

				fourth--
			}
		}
	}

	return result
}
