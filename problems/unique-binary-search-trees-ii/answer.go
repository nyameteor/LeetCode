package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (n *TreeNode) String() string {
	return treeString(n)
}

func treeString(node *TreeNode) string {
	if node == nil {
		return "nil"
	}
	left := treeString(node.Left)
	right := treeString(node.Right)
	return fmt.Sprintf("%d(%s, %s)", node.Val, left, right)
}

func generateTrees(n int) []*TreeNode {
	return buildTrees(1, n)
}

func buildTrees(low, high int) []*TreeNode {
	if low > high {
		return []*TreeNode{nil}
	}

	trees := []*TreeNode{}
	for i := low; i <= high; i++ {
		leftTrees := buildTrees(low, i-1)
		rightTrees := buildTrees(i+1, high)

		for _, left := range leftTrees {
			for _, right := range rightTrees {
				root := &TreeNode{
					Val:   i,
					Left:  left,
					Right: right,
				}
				trees = append(trees, root)
			}
		}
	}
	return trees
}

func main() {
	testCases := []struct {
		input int
	}{
		{
			input: 3,
		},
		{
			input: 1,
		},
	}

	for _, tc := range testCases {
		fmt.Println(generateTrees(tc.input))
	}
}
