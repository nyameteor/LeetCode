package main

import (
	"fmt"
	"slices"
)

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	graph := buildGraph(n, edges)
	degrees := make([]int, n)

	for node, neighbors := range graph {
		degrees[node] = len(neighbors)
	}

	queue := make([]int, 0, n)

	for node, degree := range degrees {
		if degree == 1 {
			queue = append(queue, node)
		}
	}

	// Trim leaves until we reach at most 2 nodes
	remainNodes := n
	for remainNodes > 2 {
		leafNodes := len(queue)
		remainNodes -= leafNodes

		for range leafNodes {
			node := queue[0]
			queue = queue[1:]

			for _, neighbor := range graph[node] {
				degrees[neighbor]--
				if degrees[neighbor] == 1 {
					queue = append(queue, neighbor)
				}
			}
		}
	}

	return queue
}

func buildGraph(n int, edges [][]int) [][]int {
	graph := make([][]int, n)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}
	return graph
}

func main() {
	type Input struct {
		n     int
		edges [][]int
	}

	testCases := []struct {
		input    *Input
		expected []int
	}{
		{
			input: &Input{
				n:     4,
				edges: [][]int{{1, 0}, {1, 2}, {1, 3}},
			},
			expected: []int{1},
		},
		{
			input: &Input{
				n:     6,
				edges: [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}},
			},
			expected: []int{3, 4},
		},
		{
			input: &Input{
				n:     1,
				edges: [][]int{},
			},
			expected: []int{0},
		},
	}

	for _, tc := range testCases {
		actual := findMinHeightTrees(tc.input.n, tc.input.edges)
		if !slices.Equal(actual, tc.expected) {
			fmt.Printf("findMinHeightTrees(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
