package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var list = createListNode([]int{1, 2, 3, 4, 5})
	fmt.Println(createNums(reverseList(list)))
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	prev := head
	cur := head.Next
	next := cur

	for cur != nil {
		next = cur.Next
		cur.Next = prev

		if prev == head {
			prev.Next = nil
		}

		prev = cur
		cur = next
	}

	return prev
}

func createListNode(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &ListNode{
		Val:  nums[0],
		Next: nil,
	}

	p := head
	for i := 1; i < len(nums); i++ {
		p.Next = &ListNode{
			Val:  nums[i],
			Next: nil,
		}
		p = p.Next
	}

	return head
}

func createNums(head *ListNode) []int {
	if head == nil {
		return nil
	}

	nums := []int{}
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}

	return nums
}
