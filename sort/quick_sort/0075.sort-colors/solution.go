package main

import "fmt"

func sortColors(nums []int) {
	sortWithTarget(nums, 1)
}

func sortWithTarget(nums []int, target int) {
	i, lessIdx, greaterIdx := 0, -1, len(nums)

	for i < greaterIdx {
		if nums[i] < target {
			lessIdx++
			nums[lessIdx], nums[i] = nums[i], nums[lessIdx]
			i++
		} else if nums[i] == target {
			i++
		} else if target < nums[i] {
			greaterIdx--
			nums[greaterIdx], nums[i] = nums[i], nums[greaterIdx]
		}
	}
}

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Printf("%+v\n", nums)
}
