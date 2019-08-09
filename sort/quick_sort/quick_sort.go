package main

import (
	"fmt"
)

func QuickSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}

	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, low int, high int) {
	if high <= low {
		return
	}

	i, j, pivot := low, high, arr[low]

	for i < j {
		for i < j && pivot <= arr[j] {
			j--
		}

		for i < j && arr[i] <= pivot {
			i++
		}

		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[low], arr[i] = arr[i], pivot

	quickSort(arr, low, i-1)
	quickSort(arr, i+1, high)
}

func main() {
	arr := []int{8, 2, 6, 1, 3, 4, 7, 9, 0}
	QuickSort(arr)
	fmt.Printf("%v\n", arr)
}
