package main

import "container/heap"

type IntHeap struct {
	data  []int
	isMin bool
}

func NewMinHeap(data []int) *IntHeap {
	return &IntHeap{data: data, isMin: true}
}

func NewMaxHeap(data []int) *IntHeap {
	return &IntHeap{data: data, isMin: false}
}

func (h *IntHeap) Len() int {
	return len(h.data)
}

func (h *IntHeap) Less(i int, j int) bool {
	if h.isMin {
		return h.data[i] < h.data[j]
	}
	return h.data[i] > h.data[j]
}

func (h *IntHeap) Swap(i int, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *IntHeap) Push(x any) {
	h.data = append(h.data, x.(int))
}

func (h *IntHeap) Pop() any {
	n := len(h.data)
	x := h.data[n-1]
	h.data = h.data[:n-1]
	return x
}

type MedianFinder struct {
	maxHeap *IntHeap // stores smaller half (as a max heap)
	minHeap *IntHeap // stores larger half (as a min heap)
	isEven  bool
}

func Constructor() MedianFinder {
	return MedianFinder{
		maxHeap: NewMaxHeap([]int{}),
		minHeap: NewMinHeap([]int{}),
		isEven:  true,
	}
}

func (this *MedianFinder) AddNum(num int) {
	if this.isEven {
		heap.Push(this.maxHeap, num)
		heap.Push(this.minHeap, heap.Pop(this.maxHeap))
	} else {
		heap.Push(this.minHeap, num)
		heap.Push(this.maxHeap, heap.Pop(this.minHeap))
	}
	this.isEven = !this.isEven
}

func (this *MedianFinder) FindMedian() float64 {
	if this.isEven {
		return (float64(this.minHeap.data[0]) + float64(this.maxHeap.data[0])) / 2
	} else {
		return float64(this.minHeap.data[0])
	}
}

func main() {
	medianFinder := Constructor()
	medianFinder.AddNum(1)    // arr = [1]
	medianFinder.AddNum(2)    // arr = [1, 2]
	medianFinder.FindMedian() // return 1.5 (i.e., (1 + 2) / 2)
	medianFinder.AddNum(3)    // arr[1, 2, 3]
	medianFinder.FindMedian() // return 2.0
}
