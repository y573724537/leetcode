## 思路
可以使用二分法解题

```go
func searchInsert(nums []int, target int) int {
    if len(nums) == 0 {
        return 0
    }

    low, high := 0, len(nums) - 1

    for low <= high {
        mid := (low + high) / 2

        if nums[mid] == target {
            return mid
        }

        if target < nums[mid] {
            high = mid - 1
        } else {
            low = mid + 1
        }
    }

    return low
}
```

[1,3,5,6], 2
low = 0 high = 3 mid = 1 nums[mid] = 3 => high = mid - 1 = 0