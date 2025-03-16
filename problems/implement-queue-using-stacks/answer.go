package main

import "fmt"

type Stack []int

func (s *Stack) Push(x int) {
	*s = append(*s, x)
}

func (s *Stack) Pop() int {
	n := len(*s)
	x := (*s)[n-1]
	*s = (*s)[:n-1]
	return x
}

func (s *Stack) Peek() int {
	return (*s)[len(*s)-1]
}

func (s Stack) Empty() bool {
	return len(s) == 0
}

type MyQueue struct {
	input  *Stack
	output *Stack
}

func Constructor() MyQueue {
	return MyQueue{
		input:  &Stack{},
		output: &Stack{},
	}
}

func (this *MyQueue) Push(x int) {
	this.input.Push(x)
}

func (this *MyQueue) Pop() int {
	if this.output.Empty() {
		for !this.input.Empty() {
			this.output.Push(this.input.Pop())
		}
	}
	return this.output.Pop()
}

func (this *MyQueue) Peek() int {
	if this.output.Empty() {
		for !this.input.Empty() {
			this.output.Push(this.input.Pop())
		}
	}
	return this.output.Peek()
}

func (this *MyQueue) Empty() bool {
	return this.output.Empty() && this.input.Empty()
}

func main() {
	queue := Constructor()
	queue.Push(1)
	queue.Push(2)
	// 1
	fmt.Println(queue.Peek())
	// 1
	fmt.Println(queue.Pop())
	// false
	fmt.Println(queue.Empty())
}
