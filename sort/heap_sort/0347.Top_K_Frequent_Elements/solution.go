package main

import (
	"container/heap"
	"fmt"
)

func topKFrequent(nums []int, k int) []int {
	records := make(map[int]int, 0)
	for _, num := range nums {
		records[num]++
	}

	kfHeap := NewKFHeap(k)
	heap.Init(kfHeap)
	for key, freq := range records {
		heap.Push(kfHeap, &KF{Key: key, Freq: freq})
		if k < kfHeap.Len() {
			heap.Pop(kfHeap)
		}
	}

	topKFreq := make([]int, k)
	for i, kf := range kfHeap.Data {
		topKFreq[i] = kf.Key
	}

	return topKFreq
}

type KF struct {
	Key  int
	Freq int
}

type KFHeap struct {
	Data []*KF
}

func NewKFHeap(capacity int) *KFHeap {
	return &KFHeap{
		Data: make([]*KF, 0, capacity+1),
	}
}

func (h *KFHeap) Len() int {
	return len(h.Data)
}

func (h *KFHeap) Less(i, j int) bool {
	return h.Data[i].Freq < h.Data[j].Freq
}

func (h *KFHeap) Swap(i, j int) {
	h.Data[i], h.Data[j] = h.Data[j], h.Data[i]
}

func (h *KFHeap) Push(ele interface{}) {
	h.Data = append(h.Data, ele.(*KF))
}

func (h *KFHeap) Pop() interface{} {
	ele := h.Data[h.Len()-1]
	h.Data = h.Data[:h.Len()-1]

	return ele
}

func main() {
	fmt.Printf("%v\n", topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
}
