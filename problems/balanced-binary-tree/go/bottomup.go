package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	return getHeight(root) != -1
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lh := getHeight(root.Left)
	if lh == -1 {
		return -1
	}

	rh := getHeight(root.Right)
	if rh == -1 {
		return -1
	}

	if abs(lh-rh) > 1 {
		return -1
	}

	return 1 + max(lh, rh)
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func abs(a int) int {
	if a >= 0 {
		return a
	} else {
		return -a
	}
}
