package main

import (
	"container/heap"
	"fmt"
	"maps"
)

type FreqEntry struct {
	num   int
	count int
}

type FreqHeap []*FreqEntry

func (h FreqHeap) Len() int {
	return len(h)
}

func (h FreqHeap) Less(i int, j int) bool {
	return h[i].count < h[j].count
}

func (h FreqHeap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *FreqHeap) Push(x any) {
	*h = append(*h, x.(*FreqEntry))
}

func (h *FreqHeap) Pop() any {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)

	for _, num := range nums {
		freq[num]++
	}

	h := &FreqHeap{}

	for num, count := range freq {
		heap.Push(h, &FreqEntry{num, count})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	res := []int{}
	for h.Len() > 0 {
		num := heap.Pop(h).(*FreqEntry).num
		res = append(res, num)
	}

	return res
}

func main() {
	type Input struct {
		nums []int
		k    int
	}

	testCases := []struct {
		input    *Input
		expected []int
	}{
		{
			input: &Input{
				nums: []int{1, 1, 1, 2, 2, 3},
				k:    2,
			},
			expected: []int{1, 2},
		},
		{
			input: &Input{
				nums: []int{1},
				k:    1,
			},
			expected: []int{1},
		},
	}

	setsEqual := func(s1, s2 []int) bool {
		m1, m2 := make(map[int]struct{}), make(map[int]struct{})
		for _, e := range s1 {
			m1[e] = struct{}{}
		}
		for _, s := range s2 {
			m2[s] = struct{}{}
		}
		return maps.Equal(m1, m2)
	}

	for _, tc := range testCases {
		actual := topKFrequent(tc.input.nums, tc.input.k)
		if !setsEqual(actual, tc.expected) {
			fmt.Printf("topKFrequent(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
