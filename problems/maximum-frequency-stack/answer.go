package main

import "fmt"

type FreqStack struct {
	freq   map[int]int
	stacks [][]int
}

func Constructor() FreqStack {
	return FreqStack{
		freq:   make(map[int]int),
		stacks: make([][]int, 0),
	}
}

func (s *FreqStack) Push(val int) {
	s.freq[val]++

	count := s.freq[val]
	if count > len(s.stacks) {
		s.stacks = append(s.stacks, nil)
	}
	s.stacks[count-1] = append(s.stacks[count-1], val)
}

func (s *FreqStack) Pop() int {
	stack := s.stacks[len(s.stacks)-1]
	val := stack[len(stack)-1]

	s.stacks[len(s.stacks)-1] = stack[:len(stack)-1]

	if len(s.stacks[len(s.stacks)-1]) == 0 {
		s.stacks = s.stacks[:len(s.stacks)-1]
	}

	s.freq[val]--
	return val
}

func main() {
	freqStack := Constructor()

	freqStack.Push(5) // The stack is [5]
	freqStack.Push(7) // The stack is [5,7]
	freqStack.Push(5) // The stack is [5,7,5]
	freqStack.Push(7) // The stack is [5,7,5,7]
	freqStack.Push(4) // The stack is [5,7,5,7,4]
	freqStack.Push(5) // The stack is [5,7,5,7,4,5]

	// return 5, as 5 is the most frequent. The stack becomes [5,7,5,7,4].
	fmt.Println(freqStack.Pop())

	// return 7, as 5 and 7 is the most frequent, but 7 is closest to the top. The stack becomes [5,7,5,4].
	fmt.Println(freqStack.Pop())

	// return 5, as 5 is the most frequent. The stack becomes [5,7,4].
	fmt.Println(freqStack.Pop())

	// return 4, as 4, 5 and 7 is the most frequent, but 4 is closest to the top. The stack becomes [5,7]
	fmt.Println(freqStack.Pop())
}
