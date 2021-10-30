## 思路：
### 1、暴力解法
对于给定数组每个元素，遍历整个数组，寻找另一个与之求和为给定目标的元素
```go
func twoSum(nums []int, target int) []int {
    for i := 0; i < len(nums); i++ {
        for j := 0; j < len(nums); j++ {
            if i != j && nums[i] + nums[j] == target {
                return []int{i, j}
            }
        }
    }

    return nil
}
```

缺点：时间复杂度为O(n<sup>2</sup>)

### 2、暴力解法优化
基于解法1做部分优化，寻找索引为i的配对元素时，可以从索引i+1开始寻找，理由是在寻找索引(0 ~ i - 1)时已经验证过与索引i不匹配
```go
func twoSum(nums []int, target int) []int {
    for i := 0; i < len(nums); i++ {
        for j := i + 1; j < len(nums); j++ {
            if nums[i] + nums[j] == target {
                return []int{i, j}
            }
        }
    }

    return nil
}
```

缺点：虽然有所优化，但起时间复杂度依然为O(n<sup>2</sup>)

### 3、查找解法
对于每个元素，先计算出和其匹配的对应值，在数组中查找是否有此对应值，如果有，且不是当前元素的索引，则为目标结果
```go
func twoSum(nums []int, target int) []int {
    for i := 0; i < len(nums); i++ {
        mate := target - nums[i]
        mateIdx, ok := find(nums, mate)
        if ok && mateIdx != i {
            return []int{i, mateIdx}
        }
    }

    return nil
}
```

该解法核心在于find函数，如何快速的查找到mate在nums中到索引  

如果继续遍历数组，则时间复杂度依然为O(n<sup>2</sup>)  

这里我们可以使用map数据结构，map在查找时时间复杂度为O(1)

```go
func twoSum(nums []int, target int) []int {
    records := make(map[int]int, len(nums))

    for i := 0; i < len(nums); i++ {
        mate := target - nums[i]
        if mateIdx, ok := records[mate]; ok {
            return []int{mateIdx, i}
        }

        records[nums[i]] = i
    }

    return nil
}
```

对于索引为i的元素，其匹配的对应值在记录中进行查找，如果找不到，说明索引(0 ~ i - 1)没有与其匹配的值，将索引i的元素以元素值为key，索引为value存入记录中；如果找到，则说明，索引(0 ~ i - 1)中，有与其匹配的值，且其索引不为i，则可将{mateIdx, i}作为结果返回。

该解法时间复杂度为O(n)