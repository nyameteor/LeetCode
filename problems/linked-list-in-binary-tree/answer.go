package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func checkPath(cur *ListNode, root *TreeNode) bool {
	if cur == nil {
		return true
	}

	if root == nil {
		return false
	}

	if cur.Val != root.Val {
		return false
	}

	return checkPath(cur.Next, root.Left) || checkPath(cur.Next, root.Right)
}

func isSubPath(head *ListNode, root *TreeNode) bool {
	if root == nil {
		return false
	}

	if checkPath(head, root) {
		return true
	}

	return isSubPath(head, root.Left) || isSubPath(head, root.Right)
}

func main() {
	var head *ListNode
	var root *TreeNode
	var result bool

	root = &TreeNode{1,
		&TreeNode{4,
			nil,
			&TreeNode{2,
				&TreeNode{1, nil, nil},
				nil}},
		&TreeNode{4,
			&TreeNode{2,
				&TreeNode{6, nil, nil},
				&TreeNode{8,
					&TreeNode{1, nil, nil},
					&TreeNode{3, nil, nil}}},
			nil}}

	head = &ListNode{4, &ListNode{2, &ListNode{8, nil}}}
	result = isSubPath(head, root)
	// true
	fmt.Println(result)

	head = &ListNode{4, &ListNode{2, &ListNode{6, nil}}}
	result = isSubPath(head, root)
	// true
	fmt.Println(result)

	head = &ListNode{4, &ListNode{2, &ListNode{6, &ListNode{8, nil}}}}
	result = isSubPath(head, root)
	// false
	fmt.Println(result)
}
