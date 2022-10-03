package main

import "fmt"

type MyCircularQueue struct {
	Head     *listNode
	Tail     *listNode
	Length   int
	Capacity int
}

type listNode struct {
	Val  int
	Next *listNode
}

func Constructor(k int) MyCircularQueue {
	var dummy = &listNode{
		Val:  -1,
		Next: nil,
	}

	curNode := dummy
	for i := 0; i < k; i++ {
		curNode.Next = &listNode{
			Val:  0,
			Next: nil,
		}
		curNode = curNode.Next
	}
	curNode.Next = dummy.Next

	return MyCircularQueue{
		Head:     dummy.Next,
		Tail:     dummy,
		Length:   0,
		Capacity: k,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}

	this.Tail = this.Tail.Next
	this.Tail.Val = value
	this.Length++

	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}

	tmp := this.Head
	this.Head = this.Head.Next
	this.Length--

	if this.IsEmpty() {
		this.Tail = tmp
	}

	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.Head.Val
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.Tail.Val
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.Length == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.Length == this.Capacity
}

func main() {
	var obj MyCircularQueue

	obj = Constructor(12)
	fmt.Println(obj.IsEmpty() == true)
	fmt.Println(obj.EnQueue(69) == true)
	fmt.Println(obj.DeQueue() == true)
	fmt.Println(obj.EnQueue(92) == true)
	fmt.Println(obj.EnQueue(12) == true)
	fmt.Println(obj.DeQueue() == true)
	fmt.Println(obj.IsFull() == false)
	fmt.Println(obj.Front() == 12)
}
