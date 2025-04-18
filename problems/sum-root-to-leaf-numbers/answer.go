package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	var dfs func(node *TreeNode, pathSum int) int
	dfs = func(node *TreeNode, pathSum int) int {
		if node == nil {
			return 0
		}
		pathSum = pathSum*10 + node.Val
		if node.Left == nil && node.Right == nil {
			return pathSum
		}
		return dfs(node.Left, pathSum) + dfs(node.Right, pathSum)
	}
	return dfs(root, 0)
}

func main() {
	testCases := []struct {
		input    *TreeNode
		expected int
	}{
		{
			input: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			expected: 25,
		},
		{
			input: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   9,
					Left:  &TreeNode{Val: 5},
					Right: &TreeNode{Val: 1},
				},
				Right: &TreeNode{Val: 0},
			},
			expected: 1026,
		},
	}

	for _, tc := range testCases {
		actual := sumNumbers(tc.input)
		if actual != tc.expected {
			fmt.Printf("sumNumbers(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
