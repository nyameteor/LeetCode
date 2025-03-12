package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	count := 0
	curr := head
	for curr != nil && count < k {
		curr = curr.Next
		count++
	}

	if count == k {
		newHead := reverse(head, k)
		head.Next = reverseKGroup(curr, k)
		return newHead
	}

	// If not enough nodes, return as is
	return head
}

func reverse(head *ListNode, k int) *ListNode {
	prev := (*ListNode)(nil)
	curr := head
	for k > 0 {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
		k--
	}
	return prev
}

func main() {
	type Input struct {
		head *ListNode
		k    int
	}

	testCases := []struct {
		input    *Input
		expected *ListNode
	}{
		{
			input: &Input{
				head: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
				k:    2,
			},
			expected: &ListNode{2, &ListNode{1, &ListNode{4, &ListNode{3, &ListNode{5, nil}}}}},
		},
		{
			input: &Input{
				head: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
				k:    3,
			},
			expected: &ListNode{3, &ListNode{2, &ListNode{1, &ListNode{4, &ListNode{5, nil}}}}},
		},
	}

	listsEqual := func(l1 *ListNode, l2 *ListNode) bool {
		for l1 != nil && l2 != nil {
			if l1.Val != l2.Val {
				return false
			}
			l1, l2 = l1.Next, l2.Next
		}
		return l1 == nil && l2 == nil
	}

	for _, tc := range testCases {
		actual := reverseKGroup(tc.input.head, tc.input.k)
		if !listsEqual(actual, tc.expected) {
			fmt.Printf("reverseKGroup(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
