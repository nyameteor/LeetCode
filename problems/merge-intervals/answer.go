package main

import (
	"fmt"
	"reflect"
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		last := res[len(res)-1]
		curr := intervals[i]

		if curr[0] <= last[1] {
			last[1] = max(last[1], curr[1])
		} else {
			res = append(res, curr)
		}
	}

	return res
}

func main() {
	testCases := []struct {
		input    [][]int
		expected [][]int
	}{
		{
			input: [][]int{
				{1, 3},
				{2, 6},
				{8, 10},
				{15, 18},
			},
			expected: [][]int{
				{1, 6},
				{8, 10},
				{15, 18},
			},
		},
		{
			input: [][]int{
				{1, 4},
				{4, 5},
			},
			expected: [][]int{
				{1, 5},
			},
		},
		{
			input: [][]int{
				{1, 4},
				{0, 2},
				{3, 5},
			},
			expected: [][]int{
				{0, 5},
			},
		},
	}

	for _, tc := range testCases {
		actual := merge(tc.input)
		if !reflect.DeepEqual(actual, tc.expected) {
			fmt.Printf("merge(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
