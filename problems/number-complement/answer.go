package main

import "fmt"

func findComplement(num int) int {
	mask := 0
	for tmp := num; tmp > 0; tmp >>= 1 {
		mask = mask<<1 | 1
	}
	return mask ^ num
}

func main() {
	testCases := []struct {
		input    int
		expected int
	}{
		{
			input:    5,
			expected: 2,
		},
		{
			input:    1,
			expected: 0,
		},
	}

	for _, tc := range testCases {
		actual := findComplement(tc.input)
		if actual != tc.expected {
			fmt.Printf("findComplement(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
