package main

import (
	"fmt"
	"sort"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minimumOperations(root *TreeNode) int {
	q := []*TreeNode{}
	q = append(q, root)

	var cur *TreeNode
	var res = 0
	for len(q) > 0 {
		levelLen := len(q)
		levelValues := []int{}
		for i := 0; i < levelLen; i++ {
			cur, q = q[0], q[1:]
			levelValues = append(levelValues, cur.Val)
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}

		res += getMinSwaps(levelValues)
	}

	return res
}

func getMinSwaps(sources []int) int {
	n := len(sources)

	targets := make([]int, n)
	copy(targets, sources)
	sort.Ints(targets)

	positions := map[int]int{}
	for i, v := range sources {
		positions[v] = i
	}

	res := 0
	for i := 0; i < n; i++ {
		if sources[i] == targets[i] {
			continue
		}

		// Find the current position of the target value
		j := positions[targets[i]]

		// Update positions for the swapped values
		positions[sources[i]] = j
		positions[targets[i]] = i

		// Swap the values in the sources
		sources[j], sources[i] = sources[i], sources[j]

		res++
	}
	return res
}

func main() {
	testCases := []struct {
		input    []int
		expected int
	}{
		{
			[]int{4, 3},
			1,
		},
		{
			[]int{7, 6, 8, 5},
			2,
		},
		{
			[]int{4, 3, 2, 1},
			2,
		},
	}

	for _, tc := range testCases {
		actual := getMinSwaps(tc.input)
		if actual != tc.expected {
			fmt.Printf("getMinSwaps(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
