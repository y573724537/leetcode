package main

import (
	"fmt"
)

func InsertSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}

	for i := 1; i < len(arr); i++ {
		for j := i; 0 <= j-1 && arr[j] < arr[j-1]; j-- {
			arr[j-1], arr[j] = arr[j], arr[j-1]
		}
	}
}

func main() {
	arr := []int{8, 2, 6, 1, 3, 4, 7, 9, 0}
	InsertSort(arr)
	fmt.Printf("%v\n", arr)
}
