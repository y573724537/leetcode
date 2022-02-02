package main

import (
	"fmt"
)

func BubbleSortBase(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}

	for end := len(arr); 0 < end; end-- {
		for i := 0; i+1 < end; i++ {
			if arr[i+1] < arr[i] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
	}
}

func BubbleSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}

	for end, tmpEnd := len(arr), len(arr); 0 < end; {
		tmpEnd, end = end, 0

		for i := 0; i+1 < tmpEnd; i++ {
			if arr[i+1] < arr[i] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				end = i + 1
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
	// BubbleSort(arr)
	BubbleSortBase(arr)
	fmt.Printf("%v\n", arr)
}
