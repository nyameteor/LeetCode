package main

import (
	"container/heap"
	"fmt"
	"reflect"
	"sort"
)

func kClosest(points [][]int, k int) [][]int {
	h := &MaxHeap{}

	for _, point := range points {
		heap.Push(h, point)
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	return *h
}

type MaxHeap [][]int

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i int, j int) bool {
	return distance(h[i]) > distance(h[j])
}

func (h MaxHeap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.([]int))
}

func (h *MaxHeap) Pop() any {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}

func distance(point []int) int {
	x, y := point[0], point[1]
	return x*x + y*y
}

func compareUnordered(actual, expected [][]int) bool {
	sortPoints := func(points [][]int) {
		sort.Slice(points, func(i, j int) bool {
			if points[i][0] != points[j][0] {
				return points[i][0] < points[j][0]
			}
			return points[i][1] < points[j][1]
		})
	}

	sortPoints(actual)
	sortPoints(expected)

	return reflect.DeepEqual(actual, expected)
}

func main() {
	type Input struct {
		points [][]int
		k      int
	}

	testCases := []struct {
		input    *Input
		expected [][]int
	}{
		{
			input: &Input{
				points: [][]int{{1, 3}, {-2, 2}},
				k:      1,
			},
			expected: [][]int{{-2, 2}},
		},
		{
			input: &Input{
				points: [][]int{{3, 3}, {5, -1}, {-2, 4}},
				k:      2,
			},
			expected: [][]int{{3, 3}, {-2, 4}},
		},
	}

	for _, tc := range testCases {
		actual := kClosest(tc.input.points, tc.input.k)
		if !compareUnordered(actual, tc.expected) {
			fmt.Printf("kClosest(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
