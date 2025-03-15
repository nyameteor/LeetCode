package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	curA, curB := headA, headB
	for curA != curB {
		if curA != nil {
			curA = curA.Next
		} else {
			curA = headB
		}
		if curB != nil {
			curB = curB.Next
		} else {
			curB = headA
		}
	}
	return curA
}

func main() {
	type Input struct {
		headA, headB *ListNode
	}

	intersection1 := &ListNode{8, &ListNode{4, &ListNode{5, nil}}}
	intersection2 := &ListNode{2, &ListNode{4, nil}}

	testCases := []struct {
		input    *Input
		expected *ListNode
	}{
		{
			input: &Input{
				headA: &ListNode{4, &ListNode{1, intersection1}},
				headB: &ListNode{5, &ListNode{6, &ListNode{1, intersection1}}},
			},
			expected: intersection1,
		},
		{
			input: &Input{
				headA: &ListNode{1, &ListNode{9, &ListNode{1, intersection2}}},
				headB: &ListNode{3, intersection2},
			},
			expected: intersection2,
		},
		{
			input: &Input{
				headA: &ListNode{2, &ListNode{6, &ListNode{4, nil}}},
				headB: &ListNode{1, &ListNode{5, nil}},
			},
			expected: nil,
		},
	}

	for _, tc := range testCases {
		actual := getIntersectionNode(tc.input.headA, tc.input.headB)
		if actual != tc.expected {
			fmt.Printf("getIntersectionNode(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
