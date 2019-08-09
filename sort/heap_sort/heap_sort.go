package main

import (
	"fmt"
)

func HeapSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}

	n := len(arr)
	for i := n/2 - 1; 0 <= i; i-- {
		percolateDown(arr, n, i)
	}

	for i := 0; i < n-1; i++ {
		arr[0], arr[n-1-i] = arr[n-1-i], arr[0]
		percolateDown(arr, n-1-i, 0)
	}
}

func percolateDown(arr []int, length int, i int) {
	if length <= 2*i+1 {
		return
	}

	maxPos := 2*i + 1
	maxVal := arr[maxPos]

	if 2*i+2 < length && maxVal < arr[2*i+2] {
		maxPos = 2*i + 2
		maxVal = arr[maxPos]
	}

	if arr[i] < maxVal {
		arr[i], arr[maxPos] = arr[maxPos], arr[i]
	}

	percolateDown(arr, length, maxPos)
}

func main() {
	arr := []int{8, 2, 6, 1, 3, 4, 7, 9, 0}
	HeapSort(arr)
	fmt.Printf("%v\n", arr)
}
