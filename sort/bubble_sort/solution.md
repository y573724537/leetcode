## solution

### 核心代码

冒泡排序核心代码在于，对于从0到N每个位置的元素，只要后一个元素小于前一个元素，则进行交换，这样经过一轮迭代，最大的元素可以换到最后面

``` go
for i := 0; i+1 < sentinel; i++ {
    if arr[i+1] < arr[i] {
        arr[i], arr[i+1] = arr[i+1], arr[i]
    }
}
```

### 基础版本
基于以上核心代码，可以整理出一个基础的版本，既然每一轮都核心代码都可以将最大值移到最后，那么迭代N轮，整个数组会变得有序
```go
func BubbleSort0(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}

	for sentinel := len(arr); 0 < sentinel; sentinel-- {
		for i := 0; i+1 < sentinel; i++ {
			if arr[i+1] < arr[i] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
	}
}
```
时间复杂度:O(n<sup>2</sup>); 空间复杂度:O(1)

### 改进版本1.0
还是分析核心代码，代码从左导游遍历了每个元素，如果前一个元素大于后一个元素，则进行交换，因此通过一次遍历，如果没有发生交换，我们可以推断出该数组已经有序。
```go
func BubbleSort1(nums []int) {
	if nums == nil || len(nums) < 2 {
		return
	}

	for sentinel, swapped := len(nums), true; 0 < sentinel && swapped; sentinel-- {
		swapped = false
		for i := 0; i+1 < sentinel; i++ {
			if nums[i+1] < nums[i] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
				swapped = true
			}
		}
	}
}
```
时间复杂度: 最好情况下为O(n), 最差情况下为O(n<sup>2</sup>); 空间复杂度:O(1)

### 改进版本2.0
通过观察核心代码，可以看到，如果数组已经有序，则全程不需要交换元素；并且进一步可以看到，发生交换元素位置的后面，数组一定是有序的。

因此可以整理出一个进阶版本，即，记录交换元素的位置，如果该位置已经等于0，则可以判定，该数组已经有序。

```go
func BubbleSort2(nums []int) {
	if nums == nil || len(nums) < 2 {
		return
	}

	for sentinel := len(nums); 0 < sentinel; {
		tmpSentinel := sentinel
		sentinel = 0
		for i := 0; i+1 < tmpSentinel; i++ {
			if nums[i+1] < nums[i] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
				sentinel = i + 1
			}
		}
	}
}
```
