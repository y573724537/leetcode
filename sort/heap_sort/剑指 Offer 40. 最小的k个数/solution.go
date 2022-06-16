package main

func getLeastNumbers(arr []int, k int) []int {
	for i := len(arr)>>1 - 1; 0 <= i; i-- {
		heapify(arr, len(arr), i)
	}

	for i := 0; i < k; i++ {
		arr[0], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[0]
		heapify(arr, len(arr)-1-i, 0)
	}

	return arr[len(arr)-k:]
}

func heapify(arr []int, sentinel, i int) {
	minPos := 2*i + 1

	if sentinel <= minPos {
		return
	}

	if minPos+1 < sentinel && arr[minPos+1] < arr[minPos] {
		minPos++
	}

	if arr[i] < arr[minPos] {
		return
	}

	arr[minPos], arr[i] = arr[i], arr[minPos]
	heapify(arr, sentinel, minPos)
}
