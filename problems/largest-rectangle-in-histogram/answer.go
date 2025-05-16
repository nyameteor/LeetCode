package main

import "fmt"

func largestRectangleArea(heights []int) int {
	return monoStackSinglePass(heights)
}

func monoStackTwoPass(heights []int) int {
	n := len(heights)

	// index of first shorter bar to the leftBound
	leftBound := make([]int, n)
	stack := []int{}
	for i := 0; i < n; i++ {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			leftBound[i] = -1
		} else {
			leftBound[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	// index of first shorter bar to the right
	rightBound := make([]int, n)
	stack = stack[:0]
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			rightBound[i] = n
		} else {
			rightBound[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	res := 0
	for i := range n {
		width := rightBound[i] - leftBound[i] - 1
		area := heights[i] * width
		res = max(res, area)
	}

	return res
}

func monoStackSinglePass(heights []int) int {
	// Append a 0-height bar to force clearing the stack at the end
	heights = append(heights, 0)
	stack := []int{}
	res := 0

	for i := range heights {
		for len(stack) > 0 && heights[stack[len(stack)-1]] > heights[i] {
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// Width: if stack is empty, span is i; else, span to previous smaller bar.
			width := i
			if len(stack) > 0 {
				width = i - stack[len(stack)-1] - 1
			}

			area := heights[last] * width
			res = max(res, area)
		}

		stack = append(stack, i)
	}

	return res
}

func main() {
	testCases := []struct {
		input    []int
		expected int
	}{
		{
			input:    []int{2, 1, 5, 6, 2, 3},
			expected: 10,
		},
		{
			input:    []int{2, 4},
			expected: 4,
		},
		{
			input:    []int{1, 1},
			expected: 2,
		},
		{
			input:    []int{4, 2, 0, 3, 2, 4, 3, 4},
			expected: 10,
		},
	}

	solutions := []struct {
		name string
		fn   func([]int) int
	}{
		{
			name: "Default Solution",
			fn:   largestRectangleArea,
		},
		{
			name: "Monotonic Stack (Two Pass, Compute Left/Right Bounds)",
			fn:   monoStackTwoPass,
		},
		{
			name: "Monotonic Stack (Single Pass, Inline Width Calculation)",
			fn:   monoStackSinglePass,
		},
	}

	for _, solution := range solutions {
		for _, tc := range testCases {
			actual := solution.fn(tc.input)
			if actual != tc.expected {
				fmt.Printf("[FAILED] Name: %s | Input: %v | Output: %v | Expected: %v\n",
					solution.name, tc.input, actual, tc.expected)
			}
		}
	}
}
