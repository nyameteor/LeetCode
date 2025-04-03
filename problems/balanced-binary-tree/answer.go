package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	return bottomUpOptimized(root)
}

func bottomUpOptimized(root *TreeNode) bool {
	var check func(node *TreeNode) int
	check = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		lh := check(node.Left)
		if lh == -1 {
			return -1
		}

		rh := check(node.Right)
		if rh == -1 {
			return -1
		}

		if abs(lh-rh) > 1 {
			return -1
		}

		return max(lh, rh) + 1
	}

	return check(root) != -1
}

func bottomUp(root *TreeNode) bool {
	var check func(node *TreeNode) (int, bool)
	check = func(node *TreeNode) (int, bool) {
		if node == nil {
			return 0, true
		}

		lh, lok := check(node.Left)
		rh, rok := check(node.Right)
		balanced := lok && rok && abs(lh-rh) <= 1

		return max(lh, rh) + 1, balanced
	}

	_, ok := check(root)
	return ok
}

func topDown(root *TreeNode) bool {
	var height func(node *TreeNode) int
	height = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		return 1 + max(height(node.Left), height(node.Right))
	}

	var check func(node *TreeNode) bool
	check = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		lh := height(node.Left)
		rh := height(node.Right)
		return abs(lh-rh) <= 1 && check(node.Left) && check(node.Right)
	}

	return check(root)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	testCases := []struct {
		input    *TreeNode
		expected bool
	}{
		{
			input: &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{
					Val:   20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			},
			expected: true,
		},
		{
			input: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:   3,
						Left:  &TreeNode{Val: 4},
						Right: &TreeNode{Val: 4},
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
				Right: &TreeNode{Val: 2},
			},
			expected: false,
		},
		{
			input:    nil,
			expected: true,
		},
	}

	solutions := []struct {
		name string
		fn   func(*TreeNode) bool
	}{
		{
			name: "Default Solution",
			fn:   isBalanced,
		},
		{
			name: "Bottom-up",
			fn:   bottomUp,
		},
		{
			name: "Bottom-up Optimized",
			fn:   bottomUpOptimized,
		},
		{
			name: "Top-down",
			fn:   topDown,
		},
	}

	for _, solution := range solutions {
		for _, tc := range testCases {
			actual := solution.fn(tc.input)
			if actual != tc.expected {
				fmt.Printf("[FAILED] Name: %s | Input: %v | Output: %v | Expected: %v\n",
					solution.name, tc.input, actual, tc.expected)
			}
		}
	}
}
