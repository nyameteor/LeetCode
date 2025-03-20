package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}
	if fast == nil || fast.Next == nil {
		return nil
	}
	for head != slow {
		head = head.Next
		slow = slow.Next
	}
	return head
}

func main() {
	type Input struct {
		head []int
		pos  int
	}

	testCases := []struct {
		input *Input
	}{
		{
			input: &Input{
				head: []int{3, 2, 0, -4},
				pos:  1,
			},
		},
		{
			input: &Input{
				head: []int{1, 2},
				pos:  0,
			},
		},
		{
			input: &Input{
				head: []int{1},
				pos:  -1,
			},
		},
		{
			input: &Input{
				head: []int{},
				pos:  -1,
			},
		},
	}

	buildCyclicList := func(head []int, pos int) (*ListNode, *ListNode) {
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
		return dummy.Next, cycleNode
	}

	for _, tc := range testCases {
		headNode, cycleNode := buildCyclicList(tc.input.head, tc.input.pos)
		actual := detectCycle(headNode)
		if actual != cycleNode {
			fmt.Printf("detectCycle(%v) = %v, expected = %v\n", tc.input, actual, cycleNode)
		}
	}
}
