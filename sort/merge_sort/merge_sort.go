package main

import (
	"fmt"
)

func MergeSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}

	mergeSort(arr, 0, len(arr)-1)
}

func mergeSort(arr []int, low int, high int) {
	if high <= low {
		return
	}

	mid := (low + high) / 2
	mergeSort(arr, low, mid)
	mergeSort(arr, mid+1, high)
	merge(arr, low, mid, high)
}

func merge(arr []int, low int, mid int, high int) {
	i, j, k := low, mid+1, 0
	tmpArr := make([]int, high-low+1)

	for i <= mid && j <= high {
		if arr[i] < arr[j] {
			tmpArr[k] = arr[i]
			i++
		} else {
			tmpArr[k] = arr[j]
			j++
		}
		k++
	}

	for i <= mid {
		tmpArr[k] = arr[i]
		k++
		i++
	}

	for j <= high {
		tmpArr[k] = arr[j]
		k++
		j++
	}

	copy(arr[low:], tmpArr)
}

func main() {
	arr := []int{8, 2, 6, 1, 3, 4, 7, 9, 0}
	MergeSort(arr)
	fmt.Printf("%v\n", arr)
}
