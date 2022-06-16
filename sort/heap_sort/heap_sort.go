package main

import (
	"fmt"
)

func HeapSort(arr []int) {
	n := len(arr)
	if n < 2 {
		return
	}

	for i := n/2 - 1; 0 <= i; i-- {
		heapify(arr, n, i)
	}

	for i := 0; i < n; i++ {
		arr[0], arr[n-1-i] = arr[n-1-i], arr[0]
		heapify(arr, n-1-i, 0)
	}
}

func heapify(arr []int, sentinel, i int) {
	if sentinel <= 2*i+1 {
		return
	}

	maxPos := 2*i + 1

	if maxPos+1 < sentinel && arr[maxPos] < arr[2*i+2] {
		maxPos = maxPos + 1
	}

	if arr[maxPos] <= arr[i] {
		return
	}

	arr[i], arr[maxPos] = arr[maxPos], arr[i]
	heapify(arr, sentinel, maxPos)
}

func main() {
	arr := []int{8, 2, 6, 1, 3, 4, 7, 9, 0}
	HeapSort(arr)
	fmt.Printf("%v\n", arr)
}
