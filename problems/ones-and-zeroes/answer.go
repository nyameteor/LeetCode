package main

import (
	"fmt"
)

type StateKey struct {
	m int
	n int
	i int
}

func findMaxForm(strs []string, m int, n int) int {
	memo := map[StateKey]int{}
	return helper(strs, m, n, 0, memo)
}

func helper(strs []string, m int, n int, i int, memo map[StateKey]int) int {
	key := StateKey{m, n, i}

	if v, ok := memo[key]; ok {
		return v
	}

	if i >= len(strs) {
		return 0
	}

	zeroes, ones := countZeroesAndOnes(strs[i])

	useCurrent := 0
	if m >= zeroes && n >= ones {
		useCurrent += 1 + helper(strs, m-zeroes, n-ones, i+1, memo)
	}

	skipCurrent := helper(strs, m, n, i+1, memo)

	res := max(useCurrent, skipCurrent)

	memo[key] = res

	return res
}

func countZeroesAndOnes(s string) (int, int) {
	zeroes, ones := 0, 0
	for _, c := range s {
		if c == '0' {
			zeroes++
		} else {
			ones++
		}
	}
	return zeroes, ones
}

func main() {
	type Input struct {
		strs []string
		m    int
		n    int
	}
	testCases := []struct {
		input    Input
		expected int
	}{
		{
			Input{
				strs: []string{"10", "0001", "111001", "1", "0"},
				m:    5,
				n:    3,
			},
			4,
		},
		{
			Input{
				strs: []string{"10", "0", "1"},
				m:    1,
				n:    1,
			},
			2,
		},
		{
			Input{
				strs: []string{"10", "0001", "111001", "1", "0"},
				m:    4,
				n:    3,
			},
			3,
		},
		{
			Input{
				strs: []string{"0", "0", "1", "1"},
				m:    2,
				n:    2,
			},
			4,
		},
	}

	for _, tc := range testCases {
		actual := findMaxForm(tc.input.strs, tc.input.m, tc.input.n)
		if actual != tc.expected {
			fmt.Printf("findMaxForm(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
