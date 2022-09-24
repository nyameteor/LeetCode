package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	curPath := []int{}
	resPaths := [][]int{}
	lookup(root, &targetSum, &curPath, &resPaths)

	return resPaths
}

func lookup(root *TreeNode, target *int, curPath *[]int, resPaths *[][]int) {
	if root == nil {
		return
	}

	*curPath = append(*curPath, root.Val)
	*target -= root.Val

	if *target == 0 && root.Left == nil && root.Right == nil {
		// make a copy
		resPath := make([]int, len(*curPath))
		copy(resPath, *curPath)
		*resPaths = append(*resPaths, resPath)
	}

	lookup(root.Left, target, curPath, resPaths)
	lookup(root.Right, target, curPath, resPaths)

	*target += root.Val
	*curPath = (*curPath)[:len(*curPath)-1]
}

func main() {
	var root *TreeNode
	var targetSum int

	root = &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 11,
				Left: &TreeNode{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val:   13,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   5,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
			},
		},
	}
	targetSum = 22
	fmt.Println(pathSum(root, targetSum))
}
