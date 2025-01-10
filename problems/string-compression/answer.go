package main

import (
	"fmt"
	"slices"
	"strconv"
)

func compress(chars []byte) int {
	n := len(chars)

	l := 0
	r := 0

	for r < n {
		curChar := chars[r]
		count := 0

		for r < n && chars[r] == curChar {
			r++
			count++
		}

		chars[l] = curChar
		l++

		if count > 1 {
			for _, digit := range strconv.Itoa(count) {
				chars[l] = byte(digit)
				l++
			}
		}
	}

	return l
}

func main() {
	testCases := []struct {
		input    []byte
		expected []byte
	}{
		{
			input:    []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'},
			expected: []byte{'a', '2', 'b', '2', 'c', '3'},
		},
		{
			input:    []byte{'a'},
			expected: []byte{'a'},
		},
		{
			input:    []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'},
			expected: []byte{'a', 'b', '1', '2'},
		},
		{
			input:    []byte{'a', 'a', 'a', 'b', 'b', 'a', 'a'},
			expected: []byte{'a', '3', 'b', '2', 'a', '2'},
		},
	}

	for _, tc := range testCases {
		inputCopy := slices.Clone(tc.input)
		size := compress(tc.input)
		result := tc.input[:size]
		if !slices.Equal(result, tc.expected) {
			fmt.Printf("Test failed for input: %v, Result: %v, Expected: %v\n", string(inputCopy), string(result), string(tc.expected))
		}
	}
}
