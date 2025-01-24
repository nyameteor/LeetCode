package main

import "fmt"

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	for _, req := range prerequisites {
		neighbor, node := req[0], req[1]
		graph[node] = append(graph[node], neighbor)
	}

	visited := make([]bool, numCourses)
	inStack := make([]bool, numCourses)

	for i := 0; i < numCourses; i++ {
		if !visited[i] {
			if dfsHasCycle(graph, i, visited, inStack) {
				return false
			}
		}
	}

	return true
}

func dfsHasCycle(graph [][]int, node int, visited, inStack []bool) bool {
	if inStack[node] {
		return true
	}

	if visited[node] {
		return false
	}

	inStack[node] = true

	for _, neighbor := range graph[node] {
		if dfsHasCycle(graph, neighbor, visited, inStack) {
			return true
		}
	}

	inStack[node] = false
	visited[node] = true

	return false
}

func main() {
	type Input struct {
		numCourses    int
		prerequisites [][]int
	}

	testCases := []struct {
		input    Input
		expected bool
	}{
		{
			input: Input{
				numCourses: 2,
				prerequisites: [][]int{
					{1, 0},
				},
			},
			expected: true,
		},
		{
			input: Input{
				numCourses: 2,
				prerequisites: [][]int{
					{1, 0},
					{0, 1},
				},
			},
			expected: false,
		},
	}

	for _, tc := range testCases {
		actual := canFinish(tc.input.numCourses, tc.input.prerequisites)
		if actual != tc.expected {
			fmt.Printf("canFinish(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
