## 思路
### 1、暴力解法
对于每一条竖线，遍历其他所有的竖线，并计算面积，记录其最大值作为结果返回
```go
func maxArea(height []int) int {
    maxArea := 0

    for i := 0; i < len(height); i++ {
        for j := 0; j < len(height); j++ {
            if i != j {
                area := abs(j - i) * min(height[i], height[j])
                if maxArea < area {
                    maxArea = area
                }
            }
        }
    }

    return maxArea
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func min(i, j int) int {
    if i < j {
        return i
    }
    return j
}
```

时间复杂度为O(n<sup>2</sup>)

### 2、暴力解法优化
对于第i条竖线，仅需要从第i+1条竖线开始计算并比较
```go
func maxArea(height []int) int {
    maxArea := 0

    for i := 0; i < len(height); i++ {
        for j := i + 1; j < len(height); j++ {
            area := (j - i) * min(height[i], height[j])
            if maxArea < area {
                maxArea = area
            }
        }
    }

    return maxArea
}

func min(i, j int) int {
    if i < j {
        return i
    }
    return j
}
```

时间复杂度依然为O(n<sup>2</sup>)

### 3、双指针法

```go
func maxArea(height []int) int {
    var (
        x = 0
        y = len(height) - 1
        maxArea = 0
    )

    for x < y {
        area := min(height[x], height[y]) * (y - x)
        if maxArea < area {
            maxArea = area
        }

        if height[x] <= height[y] {
            x++
        } else {
            y--
        }
    }

    return maxArea
}

func min(i, j int) int {
    if i < j {
        return i
    }
    return j
}
```