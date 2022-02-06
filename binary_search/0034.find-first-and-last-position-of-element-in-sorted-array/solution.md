## 思路

该问题涉及了二分查找两个变形问题，即查找目标左边界及右边界

### 查找左边界的二分查找

同样的我们规定查找区间为[0, len(nums)-1]
```go
func leftBound(nums []int, target int) int {
    if len(nums) == 0 {
        return -1
    }

    low, high := 0, len(nums) - 1

    for low <= high {
        mid := low + (high - low) >> 1

        if nums[mid] < target {
            low = mid + 1
        } else if target < nums[mid] {
            high = mid - 1
        } else {
            high = mid - 1
        }
    }

    if low == len(nums) || nums[low] != target {
        return -1
    }

    return low
}
```

### 查找右边界的二分查找

同样的我们规定查找区间为[0, len(nums)-1]
```go
func rightBound(nums []int, target int) int {
    if len(nums) == 0 {
        return -1
    }

    low, high := 0, len(nums) - 1

    for low <= high {
        mid := low + (high - low) >> 1

        if target < nums[mid] {
            high = mid - 1
        } else if nums[mid] < target {
            low = mid + 1
        } else {
            low = mid + 1
        }
    }

    if high < 0 || nums[high] != target {
        return -1
    }

    return high
}
```