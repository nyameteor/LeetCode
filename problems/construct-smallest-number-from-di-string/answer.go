package main

import (
	"fmt"
	"strconv"
	"strings"
)

func smallestNumber(pattern string) string {
	res := make([]int, 0)

	dfs(pattern, -1, make([]bool, 10), make([]int, 0), &res)

	return intSliceToString(res)
}

func dfs(pattern string, index int, used []bool, path []int, res *[]int) bool {
	if index >= len(pattern) {
		*res = path
		return true
	}

	// Try smaller digits first for lexicographically smallest result
	for i := 1; i <= 9; i++ {
		if used[i] {
			continue
		}

		// Prune: Ensure the pattern is satisfied
		if index >= 0 {
			prev := path[len(path)-1]
			if (pattern[index] == 'I' && prev > i) || (pattern[index] == 'D' && prev < i) {
				continue
			}
		}

		used[i] = true
		if dfs(pattern, index+1, used, append(path, i), res) {
			// Early exit when the first valid solution is found
			return true
		}
		used[i] = false
	}

	return false
}

func intSliceToString(nums []int) string {
	var sb strings.Builder
	for _, num := range nums {
		sb.WriteString(strconv.Itoa(num))
	}
	return sb.String()
}

func main() {
	testCases := []struct {
		input    string
		expected string
	}{
		{
			input:    "IIIDIDDD",
			expected: "123549876",
		},
		{
			input:    "DDD",
			expected: "4321",
		},
	}

	for _, tc := range testCases {
		actual := smallestNumber(tc.input)
		if actual != tc.expected {
			fmt.Printf("smallestNumber(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
