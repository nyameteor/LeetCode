package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) String() string {
	return stringify(n)
}

func isPalindrome(head *ListNode) bool {
	middle := findMiddle(head)
	reversed := reverse(middle)

	p1, p2 := head, reversed
	for p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1, p2 = p1.Next, p2.Next
	}
	return true
}

func findMiddle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func reverse(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var prev, curr *ListNode = nil, head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev, curr = curr, next
	}
	return prev
}

func stringify(node *ListNode) string {
	if node == nil {
		return "nil"
	}
	return fmt.Sprintf("%d(%s)", node.Val, stringify(node.Next))
}

func clone(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	return &ListNode{head.Val, clone(head.Next)}
}

func main() {
	testCases := []struct {
		input    *ListNode
		expected bool
	}{
		{
			input:    &ListNode{1, &ListNode{2, &ListNode{2, &ListNode{1, nil}}}},
			expected: true,
		},
		{
			input:    &ListNode{1, &ListNode{2, nil}},
			expected: false,
		},
	}

	for _, tc := range testCases {
		input := clone(tc.input)
		actual := isPalindrome(input)
		if actual != tc.expected {
			fmt.Printf("isPalindrome(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
