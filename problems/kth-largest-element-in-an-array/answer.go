package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

// Min-heap
func (h IntHeap) Less(i int, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	h := &IntHeap{}

	for _, num := range nums {
		heap.Push(h, num)
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	return (*h)[0]
}

func main() {
	type Input struct {
		nums []int
		k    int
	}

	testCases := []struct {
		input    *Input
		expected int
	}{
		{
			input: &Input{
				nums: []int{3, 2, 1, 5, 6, 4},
				k:    2,
			},
			expected: 5,
		},
		{
			input: &Input{
				nums: []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
				k:    4,
			},
			expected: 4,
		},
	}

	for _, tc := range testCases {
		actual := findKthLargest(tc.input.nums, tc.input.k)
		if actual != tc.expected {
			fmt.Printf("findKthLargest(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
