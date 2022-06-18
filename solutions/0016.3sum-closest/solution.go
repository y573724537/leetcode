package main

import (
	"fmt"
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	best := math.MaxInt

	for first := 0; first < len(nums); first++ {
		if 0 < first && nums[first-1] == nums[first] {
			continue
		}

		second := first + 1
		third := len(nums) - 1
		for second < third {
			if first+1 < second && nums[second-1] == nums[second] {
				second++
				continue
			}

			if third < len(nums)-1 && nums[third] == nums[third+1] {
				third--
				continue
			}

			sum := nums[first] + nums[second] + nums[third]
			if sum == target {
				return sum
			}

			if math.Abs(float64(target-sum)) < math.Abs(float64(target-best)) {
				best = sum
			}

			if sum < target {
				second++
			}

			if sum < best {
				third--
			}
		}
	}

	return best
}

func main() {
	fmt.Printf("%v", threeSumClosest([]int{1, 2, 4, 8, 16, 32, 64, 128}, 82))
}
