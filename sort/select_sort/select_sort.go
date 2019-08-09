package main

import (
	"fmt"
)

func SelectSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}

	for i := 0; i < len(arr); i++ {
		minIndex := i

		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}

		if minIndex != i {
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
		}
	}
}

func main() {
	arr := []int{8, 2, 6, 1, 3, 4, 7, 9, 0}
	SelectSort(arr)
	fmt.Printf("%v\n", arr)
}
