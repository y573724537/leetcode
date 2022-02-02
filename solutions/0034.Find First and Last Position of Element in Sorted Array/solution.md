## 思路

题目要求时间复杂度为O(log n)，可以想到使用二分查找

首先用二分查找target最早出现的位置，如果找不到这样的位置，说明数组内没有该元素，则返回[]int{-1, -1}

再用二分查找target+1最早出现的位置，该位置-1则为target最后边的位置

```go
import "sort"

func searchRange(nums []int, target int) []int {
	left := sort.SearchInts(nums, target)
	if left == len(nums) || nums[left] != target {
		return []int{-1, -1}
	}
	return []int{left, sort.SearchInts(nums, target+1) - 1}
}
```

```go
func searchRange(nums []int, target int) []int {
	left := searchMostStart(nums, target)
	if left == len(nums) || nums[left] != target {
		return []int{-1, -1}
	}
	return []int{left, searchMostStart(nums, target+1) - 1}
}

// 只要不nums[mid]不小于target，就向前找
func searchMostStart(nums []int, target int) int {
	low, high := 0, len(nums)-1

	for low <= high {
		mid := (low + high) / 2

		if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return low
}
```