package main

import (
	"fmt"
	"slices"
)

func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := make([][]int, numCourses)
	inDegrees := make([]int, numCourses)

	for _, req := range prerequisites {
		target, source := req[0], req[1]
		graph[source] = append(graph[source], target)
		inDegrees[target]++
	}

	q := make([]int, 0)

	for node := 0; node < numCourses; node++ {
		if inDegrees[node] == 0 {
			q = append(q, node)
		}
	}

	res := make([]int, 0, numCourses)

	for len(q) > 0 {
		node := q[0]
		q = q[1:]

		res = append(res, node)

		for _, neighbor := range graph[node] {
			inDegrees[neighbor]--
			if inDegrees[neighbor] == 0 {
				q = append(q, neighbor)
			}
		}
	}

	if len(res) == numCourses {
		return res
	}

	return []int{}
}

func main() {
	type Input struct {
		numCourses    int
		prerequisites [][]int
	}

	testCases := []struct {
		input       Input
		expectedSet [][]int
	}{
		{
			input: Input{
				numCourses: 2,
				prerequisites: [][]int{
					{1, 0},
				},
			},
			expectedSet: [][]int{
				{0, 1},
				{1, 0},
			},
		},
		{
			input: Input{
				numCourses: 4,
				prerequisites: [][]int{
					{1, 0},
					{2, 0},
					{3, 1},
					{3, 2},
				},
			},
			expectedSet: [][]int{
				{0, 1, 2, 3},
			},
		},
		{
			input: Input{
				numCourses:    1,
				prerequisites: [][]int{},
			},
			expectedSet: [][]int{
				{0},
			},
		},
		// Edge case: Cycle detection
		{
			input: Input{
				numCourses: 3,
				prerequisites: [][]int{
					{1, 0},
					{2, 1},
					{0, 2},
				},
			},
			expectedSet: [][]int{
				{},
			},
		},
	}

	for _, tc := range testCases {
		actual := findOrder(tc.input.numCourses, tc.input.prerequisites)

		found := false
		for _, expected := range tc.expectedSet {
			if slices.Equal(actual, expected) {
				found = true
				break
			}
		}

		if !found {
			fmt.Printf("canFinish(%v) = %v, expected one of %v\n", tc.input, actual, tc.expectedSet)
		}
	}
}
