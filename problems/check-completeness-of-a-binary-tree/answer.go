package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return false
	}

	queue := []*TreeNode{root}
	seenNull := false

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if cur == nil {
			seenNull = true
			continue
		}

		if seenNull {
			return false
		}

		queue = append(queue, cur.Left)
		queue = append(queue, cur.Right)
	}

	return true
}

func main() {
	testCases := []struct {
		input    *TreeNode
		expected bool
	}{
		{
			input: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{
					Val:  3,
					Left: &TreeNode{Val: 6},
				},
			},
			expected: true,
		},
		{
			input: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{
					Val:   3,
					Right: &TreeNode{Val: 7},
				},
			},
			expected: false,
		},
		{
			input: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 5},
					Right: &TreeNode{Val: 6},
				},
				Right: &TreeNode{Val: 3},
			},
			expected: true,
		},
		{
			input: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 6},
				},
				Right: &TreeNode{
					Val:  3,
					Left: &TreeNode{Val: 7},
				},
			},
			expected: false,
		},
	}

	for _, tc := range testCases {
		actual := isCompleteTree(tc.input)
		if actual != tc.expected {
			fmt.Printf("isCompleteTree(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
