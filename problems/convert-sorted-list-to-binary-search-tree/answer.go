package main

import (
	"fmt"
	"strings"
)

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

func (t *TreeNode) String() string {
	builder := &strings.Builder{}
	t.buildString(builder, 0)
	return builder.String()
}

func (t *TreeNode) buildString(builder *strings.Builder, depth int) {
	if t == nil {
		return
	}

	builder.WriteString(strings.Repeat(" ", depth*2))
	builder.WriteString(fmt.Sprintf("%d\n", t.Val))

	if t.Left != nil {
		t.Left.buildString(builder, depth+1)
	}
	if t.Right != nil {
		t.Right.buildString(builder, depth+1)
	}
}

func buildBST(vals []int) *TreeNode {
	if len(vals) == 0 {
		return nil
	}
	mid := len(vals) / 2
	return &TreeNode{
		Val:   vals[mid],
		Left:  buildBST(vals[:mid]),
		Right: buildBST(vals[mid+1:]),
	}
}

func sortedListToBST(head *ListNode) *TreeNode {
	vals := []int{}
	for cur := head; cur != nil; cur = cur.Next {
		vals = append(vals, cur.Val)
	}
	return buildBST(vals)
}

func main() {
	var head *ListNode
	var res *TreeNode

	head = &ListNode{-10, &ListNode{-3, &ListNode{0, &ListNode{5, &ListNode{9, nil}}}}}
	res = sortedListToBST(head)
	fmt.Println(res)

	head = nil
	res = sortedListToBST(head)
	fmt.Println(res)
}
