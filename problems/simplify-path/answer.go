package main

import (
	"fmt"
	"strings"
)

func simplifyPath(path string) string {
	n := len(path)

	stack := []string{}

	i := 0
	for i < n {
		for ; i < n && path[i] == '/'; i++ {
		}

		start := i
		for ; i < n && path[i] != '/'; i++ {
		}
		name := path[start:i]

		switch name {
		case "", ".":
		case "..":
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		default:
			stack = append(stack, name)
		}
	}

	return "/" + strings.Join(stack, "/")
}

func main() {
	testCases := []struct {
		input    string
		expected string
	}{
		{
			input:    "/home/",
			expected: "/home",
		},
		{
			input:    "/home//foo/",
			expected: "/home/foo",
		},
		{
			input:    "/home//foo/",
			expected: "/home/foo",
		},
		{
			input:    "/home/user/Documents/../Pictures",
			expected: "/home/user/Pictures",
		},
		{
			input:    "/../",
			expected: "/",
		},
		{
			input:    "/.../a/../b/c/../d/./",
			expected: "/.../b/d",
		},
		{
			input:    "/..hidden",
			expected: "/..hidden",
		},
	}

	for _, tc := range testCases {
		actual := simplifyPath(tc.input)
		if actual != tc.expected {
			fmt.Printf("simplifyPath(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
