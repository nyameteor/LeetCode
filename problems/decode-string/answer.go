package main

import (
	"fmt"
	"strings"
)

func decodeString(s string) string {
	countStack := []int{}
	strStack := []string{}
	curCount := 0
	curStr := ""

	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			curCount = curCount*10 + int(ch-'0')
		} else if ch == '[' {
			countStack = append(countStack, curCount)
			strStack = append(strStack, curStr)
			curCount = 0
			curStr = ""
		} else if ch == ']' {
			lastCount := countStack[len(countStack)-1]
			countStack = countStack[:len(countStack)-1]

			lastStr := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]

			curStr = lastStr + strings.Repeat(curStr, lastCount)
		} else {
			curStr += string(ch)
		}
	}

	return curStr
}

func main() {
	testCases := []struct {
		input    string
		expected string
	}{
		{
			input:    "3[a]2[bc]",
			expected: "aaabcbc",
		},
		{
			input:    "3[a2[c]]",
			expected: "accaccacc",
		},
		{
			input:    "2[abc]3[cd]ef",
			expected: "abcabccdcdcdef",
		},
		{
			input:    "3[z]2[2[y]pq4[2[jk]e1[f]]]ef",
			expected: "zzzyypqjkjkefjkjkefjkjkefjkjkefyypqjkjkefjkjkefjkjkefjkjkefef",
		},
	}

	for _, tc := range testCases {
		actual := decodeString(tc.input)
		if actual != tc.expected {
			fmt.Printf("decodeString(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
