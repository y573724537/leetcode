## 思路

时间复杂度要求为O(log n)，可以想到使用二分查找

由于数组有可能被rotate，因此mid的两边有可能是一边有序，一边无序
如果nums[mid] < nums[high] 那么右边一定有序
否则，左边一定有序

通过上述条件找出一定有序的一边，然后判断target是否在有序一边的范围内，如果在则将游标移动到这一边，否则将游标移动到另外一边

```go
func search(nums []int, target int) int {
    low, high := 0, len(nums) - 1

    for low <= high {
        mid := low + (high - low) / 2

        if nums[mid] == target {
            return mid
        }

        if nums[mid] < nums[high] { // 左边无序，右边有序
            if nums[mid] < target && target <= nums[high] {
                low = mid + 1
            } else {
                high = mid - 1
            }
        } else { // 左边有序，右边无序
            if target < nums[mid] && nums[low] <= target {
                high = mid - 1
            } else {
                low = mid + 1
            }
        }
    }

    return -1
}
```