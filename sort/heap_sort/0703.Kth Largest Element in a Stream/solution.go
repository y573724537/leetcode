package main

import (
	"container/heap"
	"sort"
)

type KthLargest struct {
	sort.IntSlice
	k int
}

func Constructor(k int, nums []int) KthLargest {
	instance := KthLargest{
		k: k,
	}

	for i := 0; i < len(nums); i++ {
		instance.Add(nums[i])
	}

	return instance
}

func (this *KthLargest) Push(val interface{}) {
	this.IntSlice = append(this.IntSlice, val.(int))
}

func (this *KthLargest) Pop() interface{} {
	val := this.IntSlice[len(this.IntSlice)-1]
	this.IntSlice = this.IntSlice[:len(this.IntSlice)-1]
	return val
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this, val)
	if this.Len() > this.k {
		heap.Pop(this)
	}

	return this.IntSlice[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
