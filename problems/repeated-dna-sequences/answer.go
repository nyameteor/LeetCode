package main

import (
	"fmt"
	"slices"
)

func findRepeatedDnaSequences(s string) []string {
	n := len(s)

	if n <= 10 {
		return []string{}
	}

	seqs := map[string]bool{}
	seen := map[string]bool{}

	for i := 0; i <= n-10; i++ {
		seq := s[i : i+10]
		if seen[seq] {
			seqs[seq] = true
		} else {
			seen[seq] = true
		}
	}

	res := []string{}
	for seq := range seqs {
		res = append(res, seq)
	}
	return res
}

func main() {
	testCases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT",
			expected: []string{"AAAAACCCCC", "CCCCCAAAAA"},
		},
		{
			input:    "AAAAAAAAAAAAA",
			expected: []string{"AAAAAAAAAA"},
		},
	}

	for _, tc := range testCases {
		actual := findRepeatedDnaSequences(tc.input)
		if !slices.Equal(actual, tc.expected) {
			fmt.Printf("findRepeatedDnaSequences(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
