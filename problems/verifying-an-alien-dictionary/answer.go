package main

import "fmt"

func isAlienSorted(words []string, order string) bool {
	orders := map[byte]int{}
	for i := range order {
		orders[order[i]] = i
	}

	checkOrder := func(s1, s2 string) bool {
		n := min(len(s1), len(s2))
		for i := range n {
			if orders[s1[i]] < orders[s2[i]] {
				return true
			} else if orders[s1[i]] > orders[s2[i]] {
				return false
			}
		}
		return len(s1) <= len(s2)
	}

	for i := range len(words) - 1 {
		if !checkOrder(words[i], words[i+1]) {
			return false
		}
	}
	return true
}

func main() {
	type Input struct {
		words []string
		order string
	}

	testCases := []struct {
		input    *Input
		expected bool
	}{
		{
			input: &Input{
				words: []string{"hello", "leetcode"},
				order: "hlabcdefgijkmnopqrstuvwxyz",
			},
			expected: true,
		},
		{
			input: &Input{
				words: []string{"word", "world", "row"},
				order: "worldabcefghijkmnpqstuvxyz",
			},
			expected: false,
		},
		{
			input: &Input{
				words: []string{"apple", "app"},
				order: "abcdefghijklmnopqrstuvwxyz",
			},
			expected: false,
		},
	}

	for _, tc := range testCases {
		actual := isAlienSorted(tc.input.words, tc.input.order)
		if actual != tc.expected {
			fmt.Printf("isAlienSorted(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
