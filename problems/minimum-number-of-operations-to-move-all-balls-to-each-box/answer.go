package main

import (
	"fmt"
	"slices"
)

func minOperations(boxes string) []int {
	n := len(boxes)
	res := make([]int, n)

	leftBalls := 0
	leftMoves := 0
	for i := 0; i < n; i++ {
		res[i] += leftMoves
		if boxes[i] == '1' {
			leftBalls += 1
		}
		leftMoves += leftBalls
	}

	rightBalls := 0
	rightMoves := 0
	for i := n - 1; i >= 0; i-- {
		res[i] += rightMoves
		if boxes[i] == '1' {
			rightBalls += 1
		}
		rightMoves += rightBalls
	}

	return res
}

func main() {
	testCases := []struct {
		input    string
		expected []int
	}{
		{
			input:    "110",
			expected: []int{1, 1, 3},
		},
		{
			input:    "001011",
			expected: []int{11, 8, 5, 4, 3, 4},
		},
	}

	for _, tc := range testCases {
		actual := minOperations(tc.input)
		if !slices.Equal(actual, tc.expected) {
			fmt.Printf("minOperations(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
