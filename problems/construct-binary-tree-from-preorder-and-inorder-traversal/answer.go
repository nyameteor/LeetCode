package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	rootVal := preorder[0]
	leftLen := 0
	for _, val := range inorder {
		if rootVal == val {
			break
		}
		leftLen++
	}

	root := &TreeNode{Val: rootVal}
	root.Left = buildTree(preorder[1:1+leftLen], inorder[:leftLen])
	root.Right = buildTree(preorder[1+leftLen:], inorder[leftLen+1:])
	return root
}

func main() {
	type Input struct {
		preorder []int
		inorder  []int
	}

	testCases := []struct {
		input    *Input
		expected *TreeNode
	}{
		{
			input: &Input{
				preorder: []int{3, 9, 20, 15, 7},
				inorder:  []int{9, 3, 15, 20, 7},
			},
			expected: &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{
					Val:   20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			},
		},
		{
			input: &Input{
				preorder: []int{-1},
				inorder:  []int{-1},
			},
			expected: &TreeNode{Val: -1},
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

	for _, tc := range testCases {
		actual := buildTree(tc.input.preorder, tc.input.inorder)
		if !treesEqual(actual, tc.expected) {
			fmt.Printf("buildTree(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
