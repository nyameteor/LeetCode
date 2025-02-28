package main

import "fmt"

func shortestCommonSupersequence(str1 string, str2 string) string {
	memo := make(map[[2]int]int)
	computeLCS(str1, str2, 0, 0, memo)
	res := buildSCS(str1, str2, 0, 0, memo)
	return string(res)
}

func computeLCS(str1, str2 string, i int, j int, memo map[[2]int]int) int {
	if i >= len(str1) || j >= len(str2) {
		return 0
	}

	key := [2]int{i, j}
	if val, ok := memo[key]; ok {
		return val
	}

	var res int
	if str1[i] == str2[j] {
		res = 1 + computeLCS(str1, str2, i+1, j+1, memo)
	} else {
		res = max(computeLCS(str1, str2, i+1, j, memo), computeLCS(str1, str2, i, j+1, memo))
	}

	memo[key] = res

	return res
}
func buildSCS(str1, str2 string, i, j int, memo map[[2]int]int) []byte {
	if i >= len(str1) {
		return []byte(str2[j:])
	}
	if j >= len(str2) {
		return []byte(str1[i:])
	}

	if str1[i] == str2[j] {
		return append([]byte{str1[i]}, buildSCS(str1, str2, i+1, j+1, memo)...)
	}

	if memo[[2]int{i + 1, j}] > memo[[2]int{i, j + 1}] {
		return append([]byte{str1[i]}, buildSCS(str1, str2, i+1, j, memo)...)
	} else {
		return append([]byte{str2[j]}, buildSCS(str1, str2, i, j+1, memo)...)
	}
}

func main() {
	type Input struct {
		str1 string
		str2 string
	}

	testCases := []struct {
		input    Input
		expected string
	}{
		{
			input: Input{
				str1: "abac",
				str2: "cab",
			},
			expected: "cabac",
		},
		{
			input: Input{
				str1: "aaaaaaaa",
				str2: "aaaaaaaa",
			},
			expected: "aaaaaaaa",
		},
	}

	for _, tc := range testCases {
		actual := shortestCommonSupersequence(tc.input.str1, tc.input.str2)
		if actual != tc.expected {
			fmt.Printf("shortestCommonSupersequence(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
