package main

func findKthLargest(nums []int, k int) int {
	for i := len(nums)>>1 - 1; 0 <= i; i-- {
		heapify(nums, len(nums), i)
	}

	for i := 0; i < k; i++ {
		nums[0], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[0]
		heapify(nums, len(nums)-1-i, 0)
	}

	return nums[len(nums)-k]
}

func heapify(arr []int, sentinel, i int) {
	if sentinel <= 2*i+1 {
		return
	}

	maxPos := 2*i + 1
	if maxPos+1 < sentinel && arr[maxPos] < arr[maxPos+1] {
		maxPos++
	}

	if arr[maxPos] < arr[i] {
		return
	}

	arr[maxPos], arr[i] = arr[i], arr[maxPos]

	heapify(arr, sentinel, maxPos)
}
