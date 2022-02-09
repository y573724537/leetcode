package main

import (
	"fmt"
)

func BubbleSort0(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}

	for sentinel := len(arr); 0 < sentinel; sentinel-- {
		for i := 0; i+1 < sentinel; i++ {
			if arr[i+1] < arr[i] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
	}
}

func BubbleSort1(nums []int) {
	if nums == nil || len(nums) < 2 {
		return
	}

	for sentinel, swapped := len(nums), true; 0 < sentinel && swapped; sentinel-- {
		swapped = false
		for i := 0; i+1 < sentinel; i++ {
			if nums[i+1] < nums[i] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
				swapped = true
			}
		}
	}
}

func BubbleSort2(nums []int) {
	if nums == nil || len(nums) < 2 {
		return
	}

	for sentinel := len(nums); 0 < sentinel; {
		tmpSentinel := sentinel
		sentinel = 0
		for i := 0; i+1 < tmpSentinel; i++ {
			if nums[i+1] < nums[i] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
				sentinel = i + 1
			}
		}
	}
}

func singleNumber(nums []int) []int {
	xor := 0
	for i := 0; i < len(nums); i++ {
		xor ^= nums[i]
	}
	xor ^= 0

	rightOne := xor & (^xor + 1)
	a := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]&rightOne == 1 {
			a ^= nums[i]
		}
	}
	a ^= 0

	return []int{a, xor ^ a}
}

func main() {
	arr := []int{8, 2, 6, 1, 3, 4, 7, 9, 0}
	// BubbleSort0(arr)
	// BubbleSort1(arr)
	BubbleSort2(arr)
	fmt.Printf("%v\n", arr)
}
