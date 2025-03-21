package main

import (
	"fmt"
)

func minMutation(startGene string, endGene string, bank []string) int {
	startNode, endNode := -1, -1
	for i, gene := range bank {
		if gene == startGene {
			startNode = i
		} else if gene == endGene {
			endNode = i
		}
	}

	// No valid endGene in the bank
	if endNode == -1 {
		return -1
	}

	// No valid startGene in the bank, add it at the beginning
	if startNode == -1 {
		bank = append([]string{startGene}, bank...)
		startNode, endNode = 0, endNode+1
	}

	n := len(bank)

	// Perform BFS to find the shortest path from startNode to endNode
	graph := buildGraph(bank)

	queue := []int{startNode}
	seen := make([]bool, n)
	seen[startNode] = true
	level := 0

	for len(queue) > 0 {
		size := len(queue)
		for range size {
			curNode := queue[0]
			queue = queue[1:]

			if curNode == endNode {
				return level
			}

			for _, neighbor := range graph[curNode] {
				if seen[neighbor] {
					continue
				}
				seen[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
		level++
	}

	return -1
}

func buildGraph(bank []string) [][]int {
	validMutation := func(s1, s2 string) bool {
		diff := 0
		for i := range 8 {
			if s1[i] != s2[i] {
				diff++
				if diff > 1 {
					return false
				}
			}
		}
		return diff == 1
	}

	n := len(bank)
	graph := make([][]int, n)
	for i := range n {
		for j := i + 1; j < n; j++ {
			if validMutation(bank[i], bank[j]) {
				graph[i] = append(graph[i], j)
				graph[j] = append(graph[j], i)
			}
		}
	}
	return graph
}

func main() {
	type Input struct {
		startGene string
		endGene   string
		bank      []string
	}

	testCases := []struct {
		input    *Input
		expected int
	}{
		{
			input: &Input{
				startGene: "AACCGGTT",
				endGene:   "AACCGGTA",
				bank:      []string{"AACCGGTA"},
			},
			expected: 1,
		},
		{
			input: &Input{
				startGene: "AACCGGTT",
				endGene:   "AAACGGTA",
				bank:      []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"},
			},
			expected: 2,
		},
	}

	for _, tc := range testCases {
		actual := minMutation(tc.input.startGene, tc.input.endGene, tc.input.bank)
		if actual != tc.expected {
			fmt.Printf("minMutation(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
