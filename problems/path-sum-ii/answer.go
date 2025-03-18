package main

import (
	"fmt"
	"slices"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	dfs(root, targetSum, &path, &res)
	return res
}

func dfs(root *TreeNode, remain int, path *[]int, res *[][]int) {
	if root == nil {
		return
	}

	*path = append(*path, root.Val)
	remain -= root.Val

	if root.Left == nil && root.Right == nil && remain == 0 {
		*res = append(*res, slices.Clone(*path))
	}

	dfs(root.Left, remain, path, res)
	dfs(root.Right, remain, path, res)

	*path = (*path)[:len(*path)-1]
}

func main() {
	var root *TreeNode
	var targetSum int

	root = &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val:   11,
				Left:  &TreeNode{Val: 7},
				Right: &TreeNode{Val: 2},
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val:  8,
			Left: &TreeNode{Val: 13},
			Right: &TreeNode{
				Val:   4,
				Left:  &TreeNode{Val: 5},
				Right: &TreeNode{Val: 1},
			},
		},
	}
	targetSum = 22
	fmt.Println(pathSum(root, targetSum))
}
