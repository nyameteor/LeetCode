package main

import (
	"fmt"
)

func scoreOfParentheses(s string) int {
	return splitAndEval(s)
}

func splitAndEval(s string) int {
	// Base case: "()"
	if len(s) == 2 {
		return 1
	}

	parts := splitParts(s)

	// Form: "(A)"
	if len(parts) == 1 {
		return 2 * splitAndEval(s[1:len(s)-1])
	}

	// Form: "A + B + C + ..."
	total := 0
	for i := range parts {
		total += splitAndEval(parts[i])
	}
	return total
}

// Splits string into balanced parts
func splitParts(s string) []string {
	parts := []string{}
	balance := 0
	start := 0

	for i := range s {
		if s[i] == '(' {
			balance++
		} else {
			balance--
		}
		if balance == 0 {
			parts = append(parts, s[start:i+1])
			start = i + 1
		}
	}
	return parts
}

func stackBased(s string) int {
	stack := []int{0}

	for _, c := range s {
		if c == '(' {
			stack = append(stack, 0)
		} else {
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if last == 0 {
				stack[len(stack)-1] += 1
			} else {
				stack[len(stack)-1] += last * 2
			}
		}
	}

	return stack[0]
}

func main() {
	testCases := []struct {
		input    string
		expected int
	}{
		{
			input:    "()",
			expected: 1,
		},
		{
			input:    "(())",
			expected: 2,
		},
		{
			input:    "()()",
			expected: 2,
		},
		{
			input:    "(()(()))",
			expected: 6,
		},
	}

	solutions := []struct {
		name string
		fn   func(s string) int
	}{
		{
			name: "Default Solution",
			fn:   scoreOfParentheses,
		},
		{
			name: "Split and Eval",
			fn:   splitAndEval,
		},
		{
			name: "Stack Based",
			fn:   stackBased,
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
