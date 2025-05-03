package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type pair struct {
	node *TreeNode
	pos  int
}

func widthOfBinaryTree(root *TreeNode) int {
	res := 0

	queue := []pair{{root, 1}}

	for len(queue) > 0 {
		n := len(queue)

		leftPos := queue[0].pos
		rightPos := queue[n-1].pos
		res = max(res, rightPos-leftPos+1)

		for range n {
			cur := queue[0]
			queue = queue[1:]

			if cur.node.Left != nil {
				queue = append(queue, pair{cur.node.Left, cur.pos * 2})
			}

			if cur.node.Right != nil {
				queue = append(queue, pair{cur.node.Right, cur.pos*2 + 1})
			}
		}
	}

	return res
}

func main() {
	testCases := []struct {
		input    *TreeNode
		expected int
	}{
		{
			input: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 5},
					Right: &TreeNode{Val: 3},
				},
				Right: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 9},
				},
			},
			expected: 4,
		},
		{
			input: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val:  5,
						Left: &TreeNode{Val: 6},
					},
				},
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val:  9,
						Left: &TreeNode{Val: 7},
					},
				},
			},
			expected: 7,
		},
	}

	for _, tc := range testCases {
		actual := widthOfBinaryTree(tc.input)
		if actual != tc.expected {
			fmt.Printf("widthOfBinaryTree(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
