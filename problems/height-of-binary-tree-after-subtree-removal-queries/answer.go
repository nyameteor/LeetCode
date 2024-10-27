package main

import (
	"fmt"
	"slices"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func treeQueries(root *TreeNode, queries []int) []int {
	valToHeight := map[int]int{}
	calcTreeHeight(root, valToHeight)
	rootHeight := valToHeight[root.Val]
	valToDiff := buildHeightDiffs(root, valToHeight)

	res := make([]int, len(queries))
	for i, q := range queries {
		if diff, ok := valToDiff[q]; ok {
			res[i] = rootHeight - 1 - diff
		} else {
			res[i] = rootHeight - 1
		}
	}
	return res
}

func calcTreeHeight(root *TreeNode, valToHeight map[int]int) int {
	if root == nil {
		return 0
	}
	leftHeight := calcTreeHeight(root.Left, valToHeight)
	rightHeight := calcTreeHeight(root.Right, valToHeight)
	valToHeight[root.Val] = 1 + max(leftHeight, rightHeight)
	return valToHeight[root.Val]
}

func buildHeightDiffs(root *TreeNode, valToHeight map[int]int) map[int]int {
	queue := []*TreeNode{}
	queue = append(queue, root)
	valToDiff := map[int]int{}
	var cur *TreeNode
	for len(queue) > 0 {
		curLen := len(queue)

		secondMax := 0
		firstMax := 0
		firstVal := 0
		for i := 0; i < curLen; i++ {
			cur, queue = queue[0], queue[1:]

			curHeight := valToHeight[cur.Val]
			if curHeight > firstMax {
				secondMax = firstMax
				firstMax = curHeight
				firstVal = cur.Val
			} else if curHeight > secondMax && curHeight <= firstMax {
				secondMax = curHeight
			}

			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		valToDiff[firstVal] = firstMax - secondMax
	}
	return valToDiff
}

func main() {

	type Input struct {
		root    *TreeNode
		queries []int
	}

	testCases := []struct {
		input    *Input
		expected []int
	}{
		{
			input: &Input{
				&TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 2,
						},
					},
					Right: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 6,
						},
						Right: &TreeNode{
							Val: 5,
							Right: &TreeNode{
								Val: 7,
							},
						},
					},
				},
				[]int{4},
			},
			expected: []int{2},
		},
		{
			input: &Input{
				&TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 8,
						Left: &TreeNode{
							Val: 2,
							Left: &TreeNode{
								Val: 4,
							},
							Right: &TreeNode{
								Val: 6,
							},
						},
						Right: &TreeNode{
							Val: 1,
						},
					},
					Right: &TreeNode{
						Val: 9,
						Left: &TreeNode{
							Val: 3,
						},
						Right: &TreeNode{
							Val: 7,
						},
					},
				},
				[]int{3, 2, 4, 8},
			},
			expected: []int{3, 2, 3, 2},
		},
		{
			input: &Input{
				&TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val: 5,
						Left: &TreeNode{
							Val: 3,
							Left: &TreeNode{
								Val: 2,
							},
							Right: &TreeNode{
								Val: 4,
							},
						},
					},
				},
				[]int{3, 5, 4, 2, 4},
			},
			expected: []int{1, 0, 3, 3, 3},
		},
	}

	for _, tc := range testCases {
		actual := treeQueries(tc.input.root, tc.input.queries)
		if !slices.Equal(tc.expected, actual) {
			fmt.Printf("treeQueries(%v, %v) = %v, expected = %v\n", tc.input.root, tc.input.queries, actual, tc.expected)
		}
	}
}
