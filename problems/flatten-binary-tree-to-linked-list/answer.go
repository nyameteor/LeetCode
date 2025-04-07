package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	flattenOptimized(root)
}

func flattenOptimized(root *TreeNode) {
	var prev *TreeNode

	var flattenNode func(root *TreeNode)
	flattenNode = func(root *TreeNode) {
		if root == nil {
			return
		}

		flattenNode(root.Right)
		flattenNode(root.Left)

		root.Right = prev
		root.Left = nil

		prev = root
	}

	flattenNode(root)
}

func flattenBasic(root *TreeNode) {
	if root == nil {
		return
	}

	flattenBasic(root.Left)
	flattenBasic(root.Right)

	right := root.Right

	root.Right = root.Left
	root.Left = nil

	cur := root
	for cur.Right != nil {
		cur = cur.Right
	}
	cur.Right = right
}

func main() {
	testCases := []struct {
		input    *TreeNode
		expected *TreeNode
	}{
		{
			input: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 4},
				},
				Right: &TreeNode{
					Val:   5,
					Right: &TreeNode{Val: 6},
				},
			},
			expected: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
						Right: &TreeNode{
							Val: 4,
							Right: &TreeNode{
								Val: 5,
								Right: &TreeNode{
									Val: 6,
								},
							},
						},
					},
				},
			},
		},
		{
			input:    nil,
			expected: nil,
		},
		{
			input:    &TreeNode{Val: 0},
			expected: &TreeNode{Val: 0},
		},
	}

	solutions := []struct {
		name string
		fn   func(root *TreeNode)
	}{
		{
			name: "Default Solution",
			fn:   flatten,
		},
		{
			name: "Flatten Tree Basic",
			fn:   flattenBasic,
		},
		{
			name: "Flatten Tree Optimized",
			fn:   flattenOptimized,
		},
	}

	var treesEqual func(a, b *TreeNode) bool
	treesEqual = func(a, b *TreeNode) bool {
		if a == nil && b == nil {
			return true
		}
		if a == nil || b == nil {
			return false
		}
		return a.Val == b.Val && treesEqual(a.Left, b.Left) && treesEqual(a.Right, b.Right)
	}

	var treeClone func(a *TreeNode) *TreeNode
	treeClone = func(a *TreeNode) *TreeNode {
		if a == nil {
			return nil
		}
		return &TreeNode{
			Val:   a.Val,
			Left:  treeClone(a.Left),
			Right: treeClone(a.Right),
		}
	}

	for _, solution := range solutions {
		for _, tc := range testCases {
			actual := treeClone(tc.input)
			solution.fn(actual)
			if !treesEqual(actual, tc.expected) {
				fmt.Printf("[FAILED] Name: %s | Input: %v | Output: %v | Expected: %v\n",
					solution.name, actual, actual, tc.expected)
			}
		}
	}
}
