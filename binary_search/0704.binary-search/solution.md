## 思路

此题为二分查找最基本题目，但可以根据此题目整理总结出二分查找的一些要点

```go
func search(nums []int, target int) int {
	low, high := 0, len(nums)-1

	for low <= high {
		mid := low + (high-low)>>1

		if nums[mid] == target {
			return mid
		} else if target < nums[mid] {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		}
	}

	return -1
}
```
这里由于初始化high := len(nums) - 1，决定了以下几个事情
* 所搜的区间为[low, high]
* for low <= high {}
* low = mid + 1
* high = mid - 1
* 未找到的终止状态为 low = high + 1
