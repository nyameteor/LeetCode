package main

import "fmt"

// Approach: Two pass
func canBeValid(s string, locked string) bool {
	n := len(s)

	if n%2 == 1 {
		return false
	}

	openCount := 0
	for i := 0; i < n; i++ {
		if s[i] == '(' || locked[i] == '0' {
			openCount++
		} else {
			openCount--
		}

		if openCount < 0 {
			return false
		}
	}

	closeCount := 0
	for i := n - 1; i >= 0; i-- {
		if s[i] == ')' || locked[i] == '0' {
			closeCount++
		} else {
			closeCount--
		}

		if closeCount < 0 {
			return false
		}
	}

	return true
}

// Approach: Stacks
func canBeValid2(s string, locked string) bool {
	n := len(s)

	if n%2 == 1 {
		return false
	}

	openIndices := []int{}
	unlockedIndices := []int{}

	for i := 0; i < n; i++ {
		if locked[i] == '0' {
			unlockedIndices = append(unlockedIndices, i)
		} else {
			if s[i] == '(' {
				openIndices = append(openIndices, i)
			} else {
				if len(openIndices) > 0 {
					openIndices = openIndices[:len(openIndices)-1]
				} else if len(unlockedIndices) > 0 {
					unlockedIndices = unlockedIndices[:len(unlockedIndices)-1]
				} else {
					return false
				}
			}
		}
	}

	for len(openIndices) > 0 && len(unlockedIndices) > 0 && openIndices[len(openIndices)-1] < unlockedIndices[len(unlockedIndices)-1] {
		openIndices = openIndices[:len(openIndices)-1]
		unlockedIndices = unlockedIndices[:len(unlockedIndices)-1]
	}

	if len(openIndices) == 0 && len(unlockedIndices) > 0 {
		return len(unlockedIndices)%2 == 0
	}

	return len(openIndices) == 0
}

func main() {
	type Input struct {
		s      string
		locked string
	}

	testCases := []struct {
		input    Input
		expected bool
	}{
		{
			input: Input{
				s:      "))()))",
				locked: "010100",
			},
			expected: true,
		},
		{
			input: Input{
				s:      "()()",
				locked: "0000",
			},
			expected: true,
		},
		{
			input: Input{
				s:      ")",
				locked: "0",
			},
			expected: false,
		},
		{
			input: Input{
				s:      "((()(()()))()((()()))))()((()(()",
				locked: "10111100100101001110100010001001",
			},
			expected: true,
		},
		{
			input: Input{
				s:      "())(()(()(())()())(())((())(()())((())))))(((((((())(()))))(",
				locked: "100011110110011011010111100111011101111110000101001101001111",
			},
			expected: false,
		},
	}

	for _, tc := range testCases {
		actual := canBeValid(tc.input.s, tc.input.locked)
		if actual != tc.expected {
			fmt.Printf("canBeValid(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}

		actual2 := canBeValid2(tc.input.s, tc.input.locked)
		if actual2 != tc.expected {
			fmt.Printf("canBeValid(%v) = %v, expected = %v\n", tc.input, actual2, tc.expected)
		}
	}
}
