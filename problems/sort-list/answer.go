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

func sortList(head *ListNode) *ListNode {
	return mergeSort(head)
}

func mergeSort(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return clone(head)
	}
	h1, h2 := split(head)
	h1Sorted := mergeSort(h1)
	h2Sorted := mergeSort(h2)
	return merge(h1Sorted, h2Sorted)
}

func split(h *ListNode) (*ListNode, *ListNode) {
	slow, fast := h, h
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slice(h, slow), slice(slow, nil)
}

func merge(h1, h2 *ListNode) *ListNode {
	dummy := &ListNode{-1, nil}
	tail := dummy

	for h1 != nil && h2 != nil {
		if h1.Val <= h2.Val {
			tail.Next = &ListNode{h1.Val, nil}
			tail = tail.Next
			h1 = h1.Next
		} else {
			tail.Next = &ListNode{h2.Val, nil}
			tail = tail.Next
			h2 = h2.Next
		}
	}
	if h1 != nil {
		tail.Next = clone(h1)
	}
	if h2 != nil {
		tail.Next = clone(h2)
	}

	return dummy.Next
}

func clone(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	return &ListNode{head.Val, clone(head.Next)}
}

func slice(start *ListNode, end *ListNode) *ListNode {
	if start == end {
		return nil
	}
	return &ListNode{start.Val, slice(start.Next, end)}
}

func stringify(h *ListNode) string {
	if h == nil {
		return "nil"
	}
	return fmt.Sprintf("%d(%s)", h.Val, stringify(h.Next))
}

func equal(h1, h2 *ListNode) bool {
	if h1 == nil && h2 == nil {
		return true
	}
	if h1 == nil || h2 == nil {
		return false
	}
	return h1.Val == h2.Val && equal(h1.Next, h2.Next)
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
		{
			input:    nil,
			expected: nil,
		},
	}

	for _, tc := range testCases {
		input := clone(tc.input)
		actual := sortList(input)
		if !equal(actual, tc.expected) {
			fmt.Printf("sortList(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
