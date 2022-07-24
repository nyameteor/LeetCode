package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var list []int
	var x int

	list = []int{1, 4, 3, 2, 5, 2}
	x = 3
	fmt.Println(createNums(partition(createListNode(list), x)))

	list = []int{2, 1}
	x = 2
	fmt.Println(createNums(partition(createListNode(list), x)))

	list = []int{1}
	x = 2
	fmt.Println(createNums(partition(createListNode(list), x)))

	list = []int{}
	x = 2
	fmt.Println(createNums(partition(createListNode(list), x)))
}

func partition(head *ListNode, x int) *ListNode {

	beforeHead := &ListNode{Val: 0, Next: nil}
	before := beforeHead
	afterHead := &ListNode{Val: 0, Next: nil}
	after := afterHead

	for head != nil {
		if head.Val < x {
			before.Next = head
			before = before.Next
		} else {
			after.Next = head
			after = after.Next
		}

		head = head.Next
	}

	after.Next = nil

	before.Next = afterHead.Next

	return beforeHead.Next
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
