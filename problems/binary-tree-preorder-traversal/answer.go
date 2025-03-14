package main

import (
	"fmt"
	"slices"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	left := preorderTraversal(root.Left)
	right := preorderTraversal(root.Right)

	return append(append([]int{root.Val}, left...), right...)
}

func preorderTraversalIterative(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := make([]int, 0)
	stack := []*TreeNode{root}

	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		res = append(res, cur.Val)

		// Push right first, then left (so left is processed first)
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
	}

	return res
}

func main() {
	testCases := []struct {
		input    *TreeNode
		expected []int
	}{
		{
			input: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 3},
				},
			},
			expected: []int{1, 2, 3},
		},
		{
			input: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 4},
					Right: &TreeNode{
						Val:   5,
						Left:  &TreeNode{Val: 6},
						Right: &TreeNode{Val: 7},
					},
				},
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val:  8,
						Left: &TreeNode{Val: 9},
					},
				},
			},
			expected: []int{1, 2, 4, 5, 6, 7, 3, 8, 9},
		},
		{
			input:    nil,
			expected: nil,
		},
		{
			input:    &TreeNode{Val: 1},
			expected: []int{1},
		},
	}

	type Solution struct {
		name string
		fn   func(root *TreeNode) []int
	}

	solutions := []Solution{
		{
			name: "Preorder Recursive",
			fn:   preorderTraversal,
		},
		{
			name: "Preorder Iterative",
			fn:   preorderTraversalIterative,
		},
	}

	for _, solution := range solutions {
		for _, tc := range testCases {
			actual := solution.fn(tc.input)
			if !slices.Equal(actual, tc.expected) {
				fmt.Printf("[FAILED] Name: %s | Input: %v | Output: %v | Expected: %v\n",
					solution.name, tc.input, actual, tc.expected)
			}
		}
	}
}
