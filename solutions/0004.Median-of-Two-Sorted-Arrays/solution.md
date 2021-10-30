## 思路
### 1、暴力解法
将给定两个数组合并后排序，再找出其中位数
```go
import (
    "sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    mergeNums := append(nums1, nums2...)
    sort.Ints(mergeNums)

    if len(mergeNums) % 2 != 0 {
        return float64(mergeNums[len(mergeNums) / 2])
    }

    if len(mergeNums) == 0 {
        return 0
    }

    return float64(mergeNums[len(mergeNums) / 2 - 1] + mergeNums[len(mergeNums) / 2]) / 2
}
```

该算法时间复杂度为O((m+n)log(m+n))，空间复杂度为O(m+n)

### 2、暴力解法优化
由于nums1与nums2都已经有序，因此可以使用归并的方式将数组合并

```go
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    var (
        i = 0
        j = 0
        k = 0
        mergeNums = make([]int, len(nums1) + len(nums2))
    )

    for i < len(nums1) && j < len(nums2) {
        if nums1[i] <= nums2[j] {
            mergeNums[k] = nums1[i]
            i++
        } else {
            mergeNums[k] = nums2[j]
            j++
        }
        k++
    }

    for i < len(nums1) {
        mergeNums[k] = nums1[i]
        k++
        i++
    }
    
    for j < len(nums2) {
        mergeNums[k] = nums2[j]
        k++
        j++
    }

    if len(mergeNums) % 2 != 0 {
        return float64(mergeNums[len(mergeNums) / 2])
    }

    if len(mergeNums) == 0 {
        return 0
    }

    return float64(mergeNums[len(mergeNums) / 2 - 1] + mergeNums[len(mergeNums) / 2]) / 2
}
```

该算法时间复杂度为O(m+n)，空间复杂度为O(m+n)

### 3、暴力法再优化
由于我们知道两个数组的长度，因此可以预知合并后中位数的位置，因此可以在不合并两个数组的前提下，使用指针查找到中位数。  
将问题退化为查找问题，即：寻找两个有序数组中的第K小数字。

```go
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    if (len(nums1) + len(nums2)) % 2 != 0 {
        return findKthElement(nums1, nums2, (len(nums1) + len(nums2)) / 2)
    }

    if (len(nums1) + len(nums2)) / 2 == 0 {
        return 0
    }

    return (findKthElement(nums1, nums2, (len(nums1) + len(nums2)) / 2) + findKthElement(nums1, nums2, (len(nums1) + len(nums2)) / 2 - 1)) / 2
}

func findKthElement(nums1 []int, nums2 []int, targetIdx int) float64 {
    var (
        i = 0
        j = 0
        curIdx = 0
        curV = 0
    )

    for i < len(nums1) && j < len(nums2) {
        if nums1[i] <= nums2[j] {
            curV = nums1[i]
            i++
        } else {
            curV = nums2[j]
            j++
        }

        if curIdx == targetIdx {
            return float64(curV)
        }
        curIdx++
    }

    if i == len(nums1) {
        return float64(nums2[targetIdx - len(nums1)])
    }

    return float64(nums1[targetIdx - len(nums2)])
}
```

该算法时间复杂度为O(m+n)，空间复杂度为O(1)

### 二分查找
根据上面的优化，将问题退化为查找问题，即：寻找两个有序数组中的第K小数字；若想将查找问题的时间复杂度优化为log级别，可以使用二分查找。

