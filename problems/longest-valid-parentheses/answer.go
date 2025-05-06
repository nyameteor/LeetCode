package main

import "fmt"

func longestValidParentheses(s string) int {
	return stackBased(s)
}

func stackBased(s string) int {
	res := 0
	stack := make([]int, 0, len(s))

	stack = append(stack, -1)

	for i := range s {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				res = max(res, i-stack[len(stack)-1])
			}
		}
	}

	return res
}

func topDownDp(s string) int {
	n := len(s)
	memo := make([]int, n)

	var dfs func(i int) int
	dfs = func(i int) int {
		if i < 0 {
			return 0
		}

		if memo[i] > 0 {
			return memo[i]
		}

		if s[i] == '(' {
			memo[i] = 0
		} else if i-1 >= 0 && s[i-1] == '(' {
			memo[i] = dfs(i-2) + 2
		} else if i-dfs(i-1)-1 >= 0 && s[i-dfs(i-1)-1] == '(' {
			memo[i] = dfs(i-1) + 2 + dfs(i-dfs(i-1)-2)
		}

		return memo[i]
	}

	res := 0
	for i := range n {
		res = max(res, dfs(i))
	}
	return res
}

func main() {
	testCases := []struct {
		input    string
		expected int
	}{
		{
			input:    "(()",
			expected: 2,
		},
		{
			input:    ")()())",
			expected: 4,
		},
		{
			input:    "",
			expected: 0,
		},
		{
			input:    "()(()",
			expected: 2,
		},
		{
			input:    "()",
			expected: 2,
		},
	}

	solutions := []struct {
		name string
		fn   func(s string) int
	}{
		{
			name: "Default Solution",
			fn:   longestValidParentheses,
		},
		{
			name: "Stack-Based",
			fn:   stackBased,
		},
		{
			name: "Top-down DP",
			fn:   topDownDp,
		},
	}

	for _, solution := range solutions {
		for _, tc := range testCases {
			actual := solution.fn(tc.input)
			if actual != tc.expected {
				fmt.Printf("[FAILED] Name: %s | Input: %v | Output: %v | Expected: %v\n",
					solution.name, tc.input, actual, tc.expected)
			}
		}
	}
}
