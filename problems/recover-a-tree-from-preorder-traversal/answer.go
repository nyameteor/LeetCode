package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverFromPreorder(traversal string) *TreeNode {
	n := len(traversal)

	stack := []*TreeNode{}
	i := 0

	for i < n {
		val, depth := 0, 0

		for ; traversal[i] == '-'; i++ {
			depth++
		}
		for ; i < n && traversal[i] >= '0' && traversal[i] <= '9'; i++ {
			val = val*10 + int(traversal[i]-'0')
		}

		node := &TreeNode{Val: val}

		// Adjust the stack to maintain the correct parent-child relationship
		if depth < len(stack) {
			// Pop excess nodes
			stack = stack[:depth]
		}

		// Attach the new node to its parent
		if depth > 0 {
			parent := stack[depth-1]
			if parent.Left == nil {
				parent.Left = node
			} else {
				parent.Right = node
			}
		}

		stack = append(stack, node)
	}

	return stack[0]
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

func main() {
	testCases := []struct {
		input    string
		expected *TreeNode
	}{
		{
			input: "1-2--3--4-5--6--7",
			expected: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 4},
				},
				Right: &TreeNode{
					Val:   5,
					Left:  &TreeNode{Val: 6},
					Right: &TreeNode{Val: 7},
				},
			},
		},
		{
			input: "1-2--3---4-5--6---7",
			expected: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:  3,
						Left: &TreeNode{Val: 4},
					},
				},
				Right: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val:  6,
						Left: &TreeNode{Val: 7},
					},
				},
			},
		},
		{
			input: "1-401--349---90--88",
			expected: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 401,
					Left: &TreeNode{
						Val:  349,
						Left: &TreeNode{Val: 90},
					},
					Right: &TreeNode{Val: 88},
				},
			},
		},
	}

	for _, tc := range testCases {
		actual := recoverFromPreorder(tc.input)
		if !equalTrees(actual, tc.expected) {
			fmt.Printf("recoverFromPreorder(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
