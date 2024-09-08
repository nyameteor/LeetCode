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

func (node *ListNode) String() string {
	var sb strings.Builder
	sb.WriteString("[")
	for cur := node; cur != nil; cur = cur.Next {
		sb.WriteString(fmt.Sprintf("%d", cur.Val))
		if cur.Next != nil {
			sb.WriteString(" -> ")
		}
	}
	sb.WriteString("]")
	return sb.String()
}

func splitListToParts(head *ListNode, k int) []*ListNode {
	size := 0
	for cur := head; cur != nil; cur = cur.Next {
		size += 1
	}
	splitSize := size / k
	remains := size % k
	partSize := 0

	res := []*ListNode{}
	cur := head
	for i := 0; i < k; i++ {
		partSize = splitSize
		if remains > 0 {
			partSize += 1
			remains -= 1
		}

		dummy := &ListNode{-1, nil}
		tail := dummy
		for j := 0; j < partSize; j++ {
			tail.Next = cur
			tail = tail.Next
			cur = cur.Next
		}
		tail.Next = nil
		res = append(res, dummy.Next)
	}

	return res
}

func main() {
	var head *ListNode
	var k int
	var result []*ListNode

	head = &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	k = 5
	result = splitListToParts(head, k)
	// [[1],[2],[3],[],[]]
	fmt.Println(result)

	head = &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5,
		&ListNode{6, &ListNode{7, &ListNode{8, &ListNode{9, &ListNode{10, nil}}}}}}}}}}
	k = 3
	result = splitListToParts(head, k)
	// [[1,2,3,4],[5,6,7],[8,9,10]]
	fmt.Println(result)
}
