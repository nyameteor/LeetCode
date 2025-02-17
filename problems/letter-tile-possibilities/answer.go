package main

import "fmt"

func numTilePossibilities(tiles string) int {
	n := len(tiles)

	seqs := make(map[string]struct{})
	dfs(tiles, make([]bool, n), make([]byte, 0), seqs)

	return len(seqs)
}

func dfs(tiles string, used []bool, path []byte, seqs map[string]struct{}) {
	if len(path) > 0 {
		seqs[string(path)] = struct{}{}
	}

	for i := range tiles {
		if used[i] {
			continue
		}

		used[i] = true
		dfs(tiles, used, append(path, tiles[i]), seqs)
		used[i] = false
	}
}

func main() {
	testCases := []struct {
		input    string
		expected int
	}{
		{
			input:    "AAB",
			expected: 8,
		},
		{
			input:    "AAABBC",
			expected: 188,
		},
		{
			input:    "V",
			expected: 1,
		},
	}

	for _, tc := range testCases {
		actual := numTilePossibilities(tc.input)
		if actual != tc.expected {
			fmt.Printf("numTilePossibilities(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
