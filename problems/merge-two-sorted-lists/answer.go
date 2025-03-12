package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{-1, nil}

	curr, p1, p2 := dummy, list1, list2

	for p1 != nil && p2 != nil {
		if p1.Val <= p2.Val {
			curr.Next = p1
			p1 = p1.Next
		} else {
			curr.Next = p2
			p2 = p2.Next
		}
		curr = curr.Next
	}

	if p1 != nil {
		curr.Next = p1
	} else {
		curr.Next = p2
	}

	return dummy.Next
}

func main() {
	type Input struct {
		list1 *ListNode
		list2 *ListNode
	}

	testCases := []struct {
		input    Input
		expected *ListNode
	}{
		{
			input: Input{
				list1: &ListNode{1, &ListNode{2, &ListNode{4, nil}}},
				list2: &ListNode{1, &ListNode{3, &ListNode{4, nil}}},
			},
			expected: &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{4, nil}}}}}},
		},
		{
			input: Input{
				list1: nil,
				list2: nil,
			},
			expected: nil,
		},
		{
			input: Input{
				list1: nil,
				list2: &ListNode{0, nil},
			},
			expected: &ListNode{0, nil},
		},
	}

	listsEqual := func(list1 *ListNode, list2 *ListNode) bool {
		for list1 != nil && list2 != nil {
			if list1.Val != list2.Val {
				return false
			}
			list1 = list1.Next
			list2 = list2.Next
		}
		return list1 == nil && list2 == nil
	}

	for _, tc := range testCases {
		actual := mergeTwoLists(tc.input.list1, tc.input.list2)
		if !listsEqual(actual, tc.expected) {
			fmt.Printf("mergeTwoLists(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
