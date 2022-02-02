## 思路

标准二分查找

```go
func search(nums []int, target int) int {
    low, high := 0, len(nums) - 1

    for low <= high {
        mid := (low + high) >> 1

        if nums[mid] < target {
            low = mid + 1
        } else if target < nums[mid] {
            high = mid - 1
        } else {
            return mid
        }
    }

    return -1
}
```