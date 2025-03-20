package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

func main() {
	type Input struct {
		head []int
		pos  int
	}

	testCases := []struct {
		input    *Input
		expected bool
	}{
		{
			input: &Input{
				head: []int{3, 2, 0, -4},
				pos:  1,
			},
			expected: true,
		},
		{
			input: &Input{
				head: []int{1, 2},
				pos:  0,
			},
			expected: true,
		},
		{
			input: &Input{
				head: []int{1},
				pos:  -1,
			},
			expected: false,
		},
		{
			input: &Input{
				head: []int{},
				pos:  -1,
			},
			expected: false,
		},
	}

	buildCyclicList := func(head []int, pos int) *ListNode {
		dummy := &ListNode{}
		var curNode, cycleNode *ListNode
		curNode = dummy
		for i, num := range head {
			curNode.Next = &ListNode{Val: num}
			curNode = curNode.Next
			if i == pos {
				cycleNode = curNode
			}
		}
		if pos != -1 {
			curNode.Next = cycleNode
		}
		return dummy.Next
	}

	for _, tc := range testCases {
		actual := hasCycle(buildCyclicList(tc.input.head, tc.input.pos))
		if actual != tc.expected {
			fmt.Printf("hasCycle(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
