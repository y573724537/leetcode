package solution

import "fmt"

type MaxHeap struct {
	maxHeap  []int
	length   int
	capacity int
}

func NewMaxHeap(capacity int) *MaxHeap {
	return &MaxHeap{
		maxHeap:  make([]int, 0, capacity),
		length:   0,
		capacity: capacity,
	}
}

func (heap *MaxHeap) Add(ele int) error {
	index := heap.Size()
	if index == cap(heap.maxHeap) {
		return fmt.Errorf("heap full with capacity: %v", index)
	}

	heap.maxHeap[index] = ele
	for 0 <= (index-1)>>1 && heap.maxHeap[(index-1)>>1] < heap.maxHeap[index] {
		heap.maxHeap[(index-1)>>1], heap.maxHeap[index] = heap.maxHeap[index], heap.maxHeap[(index-1)>>1]
		index = (index - 1) >> 1
	}

	heap.length++

	return nil
}

func (heap *MaxHeap) Peak() (int, error) {
	if heap.Size() == 0 {
		return 0, fmt.Errorf("heap empty")
	}

	return heap.maxHeap[0], nil
}

func (heap *MaxHeap) Pop() (int, error) {
	if heap.Size() == 0 {
		return 0, fmt.Errorf("heap empty")
	}

	top := heap.maxHeap[0]
	heap.maxHeap[0] = heap.maxHeap[heap.Size()-1]

	for i := 0; 2*i+1 < heap.Size(); {
		maxIdx := 2*i + 1
		if 2*i+2 < heap.Size() && heap.maxHeap[maxIdx] < heap.maxHeap[2*i+2] {
			maxIdx = maxIdx + 1
		}

		if heap.maxHeap[maxIdx] <= heap.maxHeap[i] {
			break
		}

		heap.maxHeap[i], heap.maxHeap[maxIdx] = heap.maxHeap[maxIdx], heap.maxHeap[i]
		i = maxIdx
	}

	return top, nil
}

func (heap *MaxHeap) Size() int {
	return heap.length
}
