package main

import "fmt"

type Node struct {
	val  int
	prev *Node
	next *Node
}

type MyCircularDeque struct {
	capacity int
	size     int
	sentinel *Node
}

func Constructor(k int) MyCircularDeque {
	sentinel := &Node{val: 0}
	sentinel.prev = sentinel
	sentinel.next = sentinel
	deque := MyCircularDeque{
		capacity: k,
		size:     0,
		sentinel: sentinel,
	}
	return deque
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	curFirst := this.sentinel.next
	newFirst := &Node{
		val:  value,
		prev: this.sentinel,
		next: curFirst,
	}
	curFirst.prev = newFirst
	this.sentinel.next = newFirst
	this.size++
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	curLast := this.sentinel.prev
	newLast := &Node{
		val:  value,
		prev: curLast,
		next: this.sentinel,
	}
	curLast.next = newLast
	this.sentinel.prev = newLast
	this.size++
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	curFirst := this.sentinel.next
	curFirst.next.prev = this.sentinel
	this.sentinel.next = curFirst.next
	this.size--
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	curLast := this.sentinel.prev
	curLast.prev.next = this.sentinel
	this.sentinel.prev = curLast.prev
	this.size--
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	first := this.sentinel.next
	return first.val
}

func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	last := this.sentinel.prev
	return last.val
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.size == 0
}

func (this *MyCircularDeque) IsFull() bool {
	return this.size == this.capacity
}

func main() {

	obj := Constructor(3)

	// true
	fmt.Println(obj.InsertLast(1))
	// true
	fmt.Println(obj.InsertLast(2))
	// true
	fmt.Println(obj.InsertFront(3))
	// false
	fmt.Println(obj.InsertFront(4))
	// 2
	fmt.Println(obj.GetRear())
	// true
	fmt.Println(obj.IsFull())
	// true
	fmt.Println(obj.DeleteLast())
	// true
	fmt.Println(obj.InsertFront(4))
	// 4
	fmt.Println(obj.GetFront())
}
