package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	_, ok := check(root)
	return ok
}

func check(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}

	lh, lok := check(root.Left)
	if !lok {
		return 0, false
	}

	rh, rok := check(root.Right)
	if !rok {
		return 0, false
	}

	if abs(lh-rh) > 1 {
		return 0, false
	}

	return 1 + max(lh, rh), true
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
