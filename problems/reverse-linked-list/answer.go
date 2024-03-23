package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	var curr *ListNode = head
	var next *ListNode

	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}

func main() {
	var list = createListNode([]int{1, 2, 3, 4, 5})
	fmt.Println(createNums(reverseList(list)))
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
