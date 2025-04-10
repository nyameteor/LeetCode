package main

import (
	"fmt"
	"math"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) String() string {
	return listString(n)
}

func listString(node *ListNode) string {
	if node == nil {
		return "nil"
	}
	return fmt.Sprintf("%d(%s)", node.Val, listString(node.Next))
}

func insertionSortList(head *ListNode) *ListNode {
	return inPlace(head)
}

func inPlace(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	sortedHead := &ListNode{math.MinInt32, head}
	curr := head

	for curr != nil && curr.Next != nil {
		// Get the target node to insert in the correct position
		target := curr.Next
		if target.Val >= curr.Val {
			curr = curr.Next
			continue
		}

		// Remove target from its current position
		curr.Next = target.Next

		// Find the correct position to insert the target node
		prev, next := sortedHead, sortedHead.Next
		for next != target && next.Val < target.Val {
			prev = next
			next = next.Next
		}

		// Insert the target node between prev and next
		prev.Next = target
		target.Next = next
	}

	return sortedHead.Next
}

func extraSpace(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	sortedHead := &ListNode{math.MinInt32, nil}
	curr := head

	for curr != nil {
		prev, next := sortedHead, sortedHead.Next

		// Find the correct position to insert the current node
		for next != nil && next.Val < curr.Val {
			prev = next
			next = next.Next
		}

		// Insert the current node into the sorted list
		prev.Next = &ListNode{curr.Val, next}

		curr = curr.Next
	}

	return sortedHead.Next
}

func main() {
	testCases := []struct {
		input    *ListNode
		expected *ListNode
	}{
		{
			input:    &ListNode{4, &ListNode{2, &ListNode{1, &ListNode{3, nil}}}},
			expected: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}},
		},
		{
			input:    &ListNode{-1, &ListNode{5, &ListNode{3, &ListNode{4, &ListNode{0, nil}}}}},
			expected: &ListNode{-1, &ListNode{0, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
		},
	}

	solutions := []struct {
		name string
		fn   func(*ListNode) *ListNode
	}{
		{
			name: "Default Solution",
			fn:   insertionSortList,
		},
		{
			name: "In-place",
			fn:   inPlace,
		},
		{
			name: "Extra Space",
			fn:   extraSpace,
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

	listClone := func(l *ListNode) *ListNode {
		dummy := &ListNode{-1, nil}
		for p := dummy; l != nil; l, p = l.Next, p.Next {
			p.Next = &ListNode{l.Val, nil}
		}
		return dummy.Next
	}

	for _, solution := range solutions {
		for _, tc := range testCases {
			input := listClone(tc.input)
			actual := solution.fn(input)
			if !listsEqual(actual, tc.expected) {
				fmt.Printf("[FAILED] Name: %s | Input: %v | Output: %v | Expected: %v\n",
					solution.name, tc.input, actual, tc.expected)
			}
		}
	}
}
