package main

import "fmt"

type Node struct {
	val    int
	minVal int
}

type MinStack struct {
	stack []Node
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	minVal := val
	if len(this.stack) > 0 {
		minVal = min(val, this.stack[len(this.stack)-1].minVal)
	}
	this.stack = append(this.stack, Node{val, minVal})
}

func (this *MinStack) Pop() {
	if len(this.stack) > 0 {
		this.stack = this.stack[:len(this.stack)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.stack) > 0 {
		last := this.stack[len(this.stack)-1]
		return last.val
	}
	return -1
}

func (this *MinStack) GetMin() int {
	if len(this.stack) > 0 {
		last := this.stack[len(this.stack)-1]
		return last.minVal
	}
	return -1
}

func main() {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Println(minStack.GetMin()) // return -3
	minStack.Pop()
	fmt.Println(minStack.Top())    // return 0
	fmt.Println(minStack.GetMin()) // return -2
}
