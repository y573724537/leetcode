package main

import (
	"fmt"
)

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

func main() {
	arr := []int{8, 2, 6, 1, 3, 4, 7, 9, 0}
	BubbleSort(arr)
	fmt.Printf("%v\n", arr)
}
