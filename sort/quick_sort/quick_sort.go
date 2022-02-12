package main

import (
	"fmt"
	"math/rand"
	"time"
)

func QuickSort0(arr []int) {
	if len(arr) < 2 {
		return
	}

	quickSort0(arr, 0, len(arr)-1)
}

func quickSort0(arr []int, low int, high int) {
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

	quickSort0(arr, low, i-1)
	quickSort0(arr, i+1, high)
}

func main() {
	arr := []int{8, 2, 6, 1, 3, 4, 7, 9, 0}
	QuickSort1(arr)
	fmt.Printf("%v\n", arr)
}

func QuickSort1(arr []int) {
	if len(arr) < 2 {
		return
	}

	quickSort1(arr, 0, len(arr)-1)
}

func quickSort1(arr []int, low, high int) {
	if high <= low {
		return
	}

	rand.Seed(time.Now().UnixNano())
	swap(arr, low+rand.Intn(high-low+1), high)

	left, right := partition(arr, low, high)

	quickSort1(arr, low, left-1)
	quickSort1(arr, right+1, high)
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func partition(arr []int, low, high int) (int, int) {
	i, leftIdx, rightIdx := low, low-1, high
	target := arr[high]

	for i < rightIdx {
		if arr[i] < target {
			leftIdx++
			swap(arr, leftIdx, i)
			i++
		} else if arr[i] == target {
			i++
		} else if target < arr[i] {
			rightIdx--
			swap(arr, rightIdx, i)
		}
	}

	swap(arr, rightIdx, high)

	return leftIdx + 1, rightIdx
}
