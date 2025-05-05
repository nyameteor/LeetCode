package main

import "fmt"

func findJudge(n int, trust [][]int) int {
	inDegrees := make([]int, n+1)
	outDegrees := make([]int, n+1)

	for _, edge := range trust {
		u, v := edge[0], edge[1]
		inDegrees[v]++
		outDegrees[u]++
	}

	for i := 1; i <= n; i++ {
		if inDegrees[i] == n-1 && outDegrees[i] == 0 {
			return i
		}
	}

	return -1
}

func main() {
	type Input struct {
		n     int
		trust [][]int
	}

	testCases := []struct {
		input    *Input
		expected int
	}{
		{
			input: &Input{
				n:     2,
				trust: [][]int{{1, 2}},
			},
			expected: 2,
		},
		{
			input: &Input{
				n:     3,
				trust: [][]int{{1, 3}, {2, 3}},
			},
			expected: 3,
		},
		{
			input: &Input{
				n:     3,
				trust: [][]int{{1, 3}, {2, 3}, {3, 1}},
			},
			expected: -1,
		},
	}

	for _, tc := range testCases {
		actual := findJudge(tc.input.n, tc.input.trust)
		if actual != tc.expected {
			fmt.Printf("findJudge(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
