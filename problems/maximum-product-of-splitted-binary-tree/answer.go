package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxProduct(root *TreeNode) int {
	totalSum := lookupSum(root)

	targetSum := 0
	lookupTargetSum(root, totalSum, &targetSum)

	res := targetSum * (totalSum - targetSum)
	return res % (1000_000_000 + 7)
}

func lookupSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return root.Val + lookupSum(root.Left) + lookupSum(root.Right)
}

// Lookup a target sum of subtree, which is the closest to the half of total sum
func lookupTargetSum(root *TreeNode, totalSum int, res *int) int {
	if root == nil {
		return 0
	}

	leftSum := lookupTargetSum(root.Left, totalSum, res)
	rightSum := lookupTargetSum(root.Right, totalSum, res)

	if abs(totalSum/2-leftSum) < abs(totalSum/2-*res) {
		*res = leftSum
	}
	if abs(totalSum/2-rightSum) < abs(totalSum/2-*res) {
		*res = rightSum
	}

	return root.Val + leftSum + rightSum
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
