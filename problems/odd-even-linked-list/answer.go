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

func stringify(node *ListNode) string {
	if node == nil {
		return "nil"
	}
	return fmt.Sprintf("%d(%s)", node.Val, stringify(node.Next))
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	odd, even := head, head.Next
	evenHead := even

	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next

		even.Next = odd.Next
		even = even.Next
	}

	odd.Next = evenHead

	return head
}

func main() {
	testCases := []struct {
		input    *ListNode
		expected *ListNode
	}{
		{
			input:    &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
			expected: &ListNode{1, &ListNode{3, &ListNode{5, &ListNode{2, &ListNode{4, nil}}}}},
		},
		{
			input:    &ListNode{2, &ListNode{1, &ListNode{3, &ListNode{5, &ListNode{6, &ListNode{4, &ListNode{7, nil}}}}}}},
			expected: &ListNode{2, &ListNode{3, &ListNode{6, &ListNode{7, &ListNode{1, &ListNode{5, &ListNode{4, nil}}}}}}},
		},
	}

	listsEqual := func(h1 *ListNode, h2 *ListNode) bool {
		for h1 != nil && h2 != nil {
			if h1.Val != h2.Val {
				return false
			}
			h1, h2 = h1.Next, h2.Next
		}
		return h1 == nil && h2 == nil
	}

	listClone := func(h *ListNode) *ListNode {
		dummy := &ListNode{-1, nil}
		for p := dummy; h != nil; h, p = h.Next, p.Next {
			p.Next = &ListNode{h.Val, nil}
		}
		return dummy.Next
	}

	for _, tc := range testCases {
		input := listClone(tc.input)
		actual := oddEvenList(input)
		if !listsEqual(actual, tc.expected) {
			fmt.Printf("oddEvenList(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
