package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	result := make([][]int, 0)

	for first := 0; first < len(nums); first++ {
		if 0 < first && first < len(nums) && nums[first-1] == nums[first] {
			continue
		}

		target := -nums[first]
		for second, third := first+1, len(nums)-1; second < third; second++ {
			if first+1 < second && second < third && nums[second-1] == nums[second] {
				continue
			}

			for second < third && target < nums[second]+nums[third] {
				third--
			}

			if second == third {
				break
			}

			if target == nums[second]+nums[third] {
				result = append(result, []int{nums[first], nums[second], nums[third]})
			}
		}

	}

	return result
}

func main() {
	fmt.Printf("%v", threeSum([]int{-1, 0, 1, 2, -1, -4}))
}
