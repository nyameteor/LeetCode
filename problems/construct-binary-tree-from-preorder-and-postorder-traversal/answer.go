package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	n := len(preorder)

	if n == 0 {
		return nil
	}

	if n == 1 {
		return &TreeNode{Val: preorder[0]}
	}

	leftChildRoot := preorder[1]
	leftChildSize := 1
	for i := 0; postorder[i] != leftChildRoot && i < n-1; i++ {
		leftChildSize++
	}

	return &TreeNode{
		Val:   preorder[0],
		Left:  constructFromPrePost(preorder[1:1+leftChildSize], postorder[:leftChildSize]),
		Right: constructFromPrePost(preorder[1+leftChildSize:], postorder[leftChildSize:n-1]),
	}
}

func equalTrees(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}
	return root1.Val == root2.Val &&
		equalTrees(root1.Left, root2.Left) &&
		equalTrees(root1.Right, root2.Right)
}

func (n *TreeNode) String() string {
	var res []string
	queue := []*TreeNode{n}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			res = append(res, "null")
			continue
		}

		res = append(res, strconv.Itoa(node.Val))
		queue = append(queue, node.Left, node.Right)
	}

	// Remove trailing "null" values to keep it compact
	for len(res) > 0 && res[len(res)-1] == "null" {
		res = res[:len(res)-1]
	}

	return "[" + strings.Join(res, ",") + "]"
}

func main() {
	type Input struct {
		preorder  []int
		postorder []int
	}

	testCases := []struct {
		input    Input
		expected *TreeNode
	}{
		{
			input: Input{
				preorder:  []int{1, 2, 4, 5, 3, 6, 7},
				postorder: []int{4, 5, 2, 6, 7, 3, 1},
			},
			expected: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 6},
					Right: &TreeNode{Val: 7},
				},
			},
		},
		{
			input: Input{
				preorder:  []int{1, 2, 3, 4},
				postorder: []int{3, 4, 2, 1},
			},
			expected: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 4},
				},
			},
		},
		{
			input: Input{
				preorder:  []int{1},
				postorder: []int{1},
			},
			expected: &TreeNode{Val: 1},
		},
	}

	for _, tc := range testCases {
		actual := constructFromPrePost(tc.input.preorder, tc.input.postorder)
		if !equalTrees(actual, tc.expected) {
			fmt.Printf("constructFromPrePost(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
