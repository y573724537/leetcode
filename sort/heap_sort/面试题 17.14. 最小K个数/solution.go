package main

func smallestK(arr []int, k int) []int {
	for i := k>>1 - 1; 0 <= i; i-- {
		heapify(arr, k, i)
	}

	for i := k; i < len(arr); i++ {
		if arr[0] < arr[i] {
			continue
		}

		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, k, 0)
	}

	return arr[:k]
}

func heapify(arr []int, sentinel, i int) {
	maxPos := 2*i + 1

	if sentinel <= maxPos {
		return
	}

	if maxPos+1 < sentinel && arr[maxPos] < arr[maxPos+1] {
		maxPos++
	}

	if arr[maxPos] < arr[i] {
		return
	}

	arr[i], arr[maxPos] = arr[maxPos], arr[i]

	heapify(arr, sentinel, maxPos)
}
