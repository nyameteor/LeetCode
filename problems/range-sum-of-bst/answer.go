package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}

	res := 0
	if root.Val >= low && root.Val <= high {
		res += root.Val
	}
	if root.Val > low {
		res += rangeSumBST(root.Left, low, high)
	}
	if root.Val < high {
		res += rangeSumBST(root.Right, low, high)
	}
	return res
}

func main() {
	type Input struct {
		root      *TreeNode
		low, high int
	}

	testCases := []struct {
		input    *Input
		expected int
	}{
		{
			input: &Input{
				root: &TreeNode{
					Val: 10,
					Left: &TreeNode{
						Val:   5,
						Left:  &TreeNode{Val: 3},
						Right: &TreeNode{Val: 7},
					},
					Right: &TreeNode{
						Val:   15,
						Right: &TreeNode{Val: 18},
					},
				},
				low:  7,
				high: 15,
			},
			expected: 32,
		},
		{
			input: &Input{
				root: &TreeNode{
					Val: 10,
					Left: &TreeNode{
						Val: 5,
						Left: &TreeNode{
							Val:  3,
							Left: &TreeNode{Val: 1},
						},
						Right: &TreeNode{
							Val:  7,
							Left: &TreeNode{Val: 6},
						},
					},
					Right: &TreeNode{
						Val:   15,
						Left:  &TreeNode{Val: 13},
						Right: &TreeNode{Val: 18},
					},
				},
				low:  6,
				high: 10,
			},
			expected: 23,
		},
	}

	for _, tc := range testCases {
		actual := rangeSumBST(tc.input.root, tc.input.low, tc.input.high)
		if actual != tc.expected {
			fmt.Printf("rangeSumBST(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
