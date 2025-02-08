package main

import (
	"fmt"
	"slices"
)

func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	graph := make([][]int, numCourses)
	inDegrees := make([]int, numCourses)
	for _, req := range prerequisites {
		source, target := req[0], req[1]
		graph[source] = append(graph[source], target)
		inDegrees[target]++
	}

	q := make([]int, 0)
	for node, inDegree := range inDegrees {
		if inDegree == 0 {
			q = append(q, node)
		}
	}

	isReachable := make([][]bool, numCourses)
	for i := 0; i < numCourses; i++ {
		isReachable[i] = make([]bool, numCourses)
	}

	for len(q) > 0 {
		node := q[0]
		q = q[1:]

		for _, neighbor := range graph[node] {
			isReachable[node][neighbor] = true

			// Propagate reachability from `node` to `neighbor`
			for i := 0; i < numCourses; i++ {
				if isReachable[i][node] {
					isReachable[i][neighbor] = true
				}
			}

			inDegrees[neighbor]--
			if inDegrees[neighbor] == 0 {
				q = append(q, neighbor)
			}
		}
	}

	res := make([]bool, len(queries))
	for i, query := range queries {
		u, v := query[0], query[1]
		res[i] = isReachable[u][v]
	}

	return res
}

func main() {
	type Input struct {
		numCourses    int
		prerequisites [][]int
		queries       [][]int
	}

	testCases := []struct {
		input    Input
		expected []bool
	}{
		{
			input: Input{
				numCourses:    2,
				prerequisites: [][]int{{1, 0}},
				queries:       [][]int{{0, 1}, {1, 0}},
			},
			expected: []bool{false, true},
		},
		{
			input: Input{
				numCourses:    2,
				prerequisites: [][]int{},
				queries:       [][]int{{0, 1}, {1, 0}},
			},
			expected: []bool{false, false},
		},
		{
			input: Input{
				numCourses:    3,
				prerequisites: [][]int{{1, 2}, {1, 0}, {2, 0}},
				queries:       [][]int{{1, 0}, {1, 2}},
			},
			expected: []bool{true, true},
		},
	}

	for _, tc := range testCases {
		actual := checkIfPrerequisite(tc.input.numCourses, tc.input.prerequisites, tc.input.queries)
		if !slices.Equal(actual, tc.expected) {
			fmt.Printf("checkIfPrerequisite(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
