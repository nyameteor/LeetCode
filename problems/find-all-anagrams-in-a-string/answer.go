package main

import (
	"fmt"
	"slices"
)

func findAnagrams(s string, p string) []int {
	res := []int{}

	sLen, pLen := len(s), len(p)

	if sLen < pLen {
		return res
	}

	var pFreq, windowFreq [26]int

	for i := range pLen {
		pFreq[p[i]-'a']++
		windowFreq[s[i]-'a']++
	}

	matchCount := 0

	for i := range 26 {
		if pFreq[i] == windowFreq[i] {
			matchCount++
		}
	}

	if matchCount == 26 {
		res = append(res, 0)
	}

	// Sliding Window
	for i := 0; i+pLen < sLen; i++ {
		left, right := s[i]-'a', s[i+pLen]-'a'

		// Remove left character
		if pFreq[left] == windowFreq[left] {
			matchCount--
		}
		windowFreq[left]--
		if pFreq[left] == windowFreq[left] {
			matchCount++
		}

		// Add right character
		if pFreq[right] == windowFreq[right] {
			matchCount--
		}
		windowFreq[right]++
		if pFreq[right] == windowFreq[right] {
			matchCount++
		}

		if matchCount == 26 {
			res = append(res, i+1)
		}
	}

	return res
}

func main() {
	type Input struct {
		s string
		p string
	}

	testCases := []struct {
		input    *Input
		expected []int
	}{
		{
			input: &Input{
				s: "cbaebabacd",
				p: "abc",
			},
			expected: []int{0, 6},
		},
		{
			input: &Input{
				s: "abab",
				p: "ab",
			},
			expected: []int{0, 1, 2},
		},
		{
			input: &Input{
				s: "aaaaaaaaaa",
				p: "aaaaaaaaaaaaa",
			},
			expected: []int{},
		},
		{
			input: &Input{
				s: "baa",
				p: "aa",
			},
			expected: []int{1},
		},
	}

	for _, tc := range testCases {
		actual := findAnagrams(tc.input.s, tc.input.p)
		if !slices.Equal(actual, tc.expected) {
			fmt.Printf("findAnagrams(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
