|堆的用法|时间复杂度|空间复杂度|
|---|---|---|
|创建堆|O(N)|O(N)|
|插入元素|O(logN)|O(1)|
|获取堆顶元素|O(1)|O(1)|
|删除堆顶元素|O(logN)|O(1)|
|获取堆的长度|O(1)|O(1)|

## golang标准库heap用法

1. 引入heap包
```go
import "container/heap"
```
2. 使用heap包的集合必须要实现heap.Interface，即需要实现如下接口
```go
type Interface interface {
    Len() int
    Less(i, j int) bool
    Swap(i, j int)
    Push(ele interface{})
    Pop() interface{}
}
```

3. heap包内实现了一个小顶堆

4. 排序时首先使用
```go
heap.Init(container)
```
完成初始化，需要注意container的长度，初始化的时候使用了container的长度进行堆化操作
如果初始时没有元素，一定记得container.Len()要返回0

5. 插入使用
```go
heap.Push(container, element)
```
对于top K类似的问题，在Push后判断container.Len()如果大于k，需要使用
```go
heap.Pop(container)
```
