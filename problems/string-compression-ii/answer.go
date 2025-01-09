package main

import (
	"fmt"
	"math"
)

func getLengthOfOptimalCompression(s string, k int) int {
	n := len(s)
	memo := make([][]int, n+1)

	for i := 0; i < n+1; i++ {
		memo[i] = make([]int, k+1)
		for j := 0; j < k+1; j++ {
			memo[i][j] = -1
		}
	}

	return helper(s, 0, k, memo)
}

func helper(s string, i int, k int, memo [][]int) int {
	n := len(s)

	if k < 0 {
		return n
	}

	if i >= n {
		return 0
	}

	if n-i <= k {
		return 0
	}

	if memo[i][k] != -1 {
		return memo[i][k]
	}

	counts := make([]int, 26)
	maxCount := 0

	res := math.MaxInt32

	for j := i; j < n; j++ {
		char := s[j] - 'a'
		counts[char]++
		maxCount = max(maxCount, counts[char])

		cost := (j - i + 1) - maxCount

		res = min(res, 1+getDigitCount(maxCount)+helper(s, j+1, k-cost, memo))
	}

	memo[i][k] = res
	return res
}

func getDigitCount(size int) int {
	if size == 1 {
		return 0
	} else if size < 10 {
		return 1
	} else if size < 100 {
		return 2
	} else {
		return 3
	}
}

func main() {
	type Input struct {
		s string
		k int
	}

	testCases := []struct {
		input    Input
		expected int
	}{
		{
			input: Input{
				s: "aaabcccd",
				k: 2,
			},
			expected: 4,
		},
		{
			input: Input{
				s: "aabbaa",
				k: 2,
			},
			expected: 2,
		},
		{
			input: Input{
				s: "aaaaaaaaaaa",
				k: 0,
			},
			expected: 3,
		},
		{
			input: Input{
				s: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
				k: 0,
			},
			expected: 4,
		},
	}

	for _, tc := range testCases {
		actual := getLengthOfOptimalCompression(tc.input.s, tc.input.k)
		if actual != tc.expected {
			fmt.Printf("getLengthOfOptimalCompression(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
