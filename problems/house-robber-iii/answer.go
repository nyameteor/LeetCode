package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	return bottomUp(root)
}

// Postorder traversal approach (bottom-up DP, no memoization)
func bottomUp(root *TreeNode) int {
	var dfs func(node *TreeNode) (rob, notRob int)
	dfs = func(node *TreeNode) (rob, notRob int) {
		if node == nil {
			return 0, 0
		}

		leftRob, leftNotRob := dfs(node.Left)
		rightRob, rightNotRob := dfs(node.Right)

		rob = node.Val + leftNotRob + rightNotRob
		notRob = max(leftRob, leftNotRob) + max(rightRob, rightNotRob)

		return rob, notRob
	}

	return max(dfs(root))
}

// Recursive approach (top-down DP with memoization)
func topDown(root *TreeNode) int {
	type Key struct {
		node       *TreeNode
		prevRobbed bool
	}

	memo := make(map[Key]int)

	var dfs func(node *TreeNode, prevRobbed bool) int
	dfs = func(node *TreeNode, prevRobbed bool) int {
		if node == nil {
			return 0
		}

		key := Key{node, prevRobbed}
		if val, ok := memo[key]; ok {
			return val
		}

		if prevRobbed {
			memo[key] = dfs(node.Left, false) + dfs(node.Right, false)
		} else {
			memo[key] = max(
				dfs(node.Left, false)+dfs(node.Right, false),
				node.Val+dfs(node.Left, true)+dfs(node.Right, true),
			)
		}

		return memo[key]
	}

	return max(dfs(root, false), dfs(root, true))
}

func main() {
	testCases := []struct {
		input    *TreeNode
		expected int
	}{
		{
			input: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 3},
				},
				Right: &TreeNode{
					Val:   3,
					Right: &TreeNode{Val: 1},
				},
			},
			expected: 7,
		},
		{
			input: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   4,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 3},
				},
				Right: &TreeNode{
					Val:   5,
					Right: &TreeNode{Val: 1},
				},
			},
			expected: 9,
		},
		{
			input: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:  2,
						Left: &TreeNode{Val: 3},
					},
				},
			},
			expected: 7,
		},
	}

	solutions := []struct {
		name string
		fn   func(root *TreeNode) int
	}{
		{
			name: "Default Solution",
			fn:   rob,
		},
		{
			name: "Top-down",
			fn:   topDown,
		},
		{
			name: "Bottom-up",
			fn:   bottomUp,
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
