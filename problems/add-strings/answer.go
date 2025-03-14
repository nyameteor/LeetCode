package main

import (
	"fmt"
)

func addStrings(num1 string, num2 string) string {
	res := make([]byte, 0)

	i, j, carry := len(num1)-1, len(num2)-1, 0
	for i >= 0 || j >= 0 || carry > 0 {
		x1, x2 := 0, 0
		if i >= 0 {
			x1 = int(num1[i] - '0')
			i--
		}
		if j >= 0 {
			x2 = int(num2[j] - '0')
			j--
		}

		sum := x1 + x2 + carry
		carry = sum / 10
		res = append(res, byte(sum%10+'0'))
	}

	reverse(res)

	return string(res)
}

func reverse[T any](s []T) {
	l, r := 0, len(s)-1
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}

func main() {
	type Input struct {
		num1, num2 string
	}

	testCases := []struct {
		input    *Input
		expected string
	}{
		{
			input: &Input{
				num1: "11",
				num2: "123",
			},
			expected: "134",
		},
		{
			input: &Input{
				num1: "456",
				num2: "77",
			},
			expected: "533",
		},
		{
			input: &Input{
				num1: "0",
				num2: "0",
			},
			expected: "0",
		},
		{
			input: &Input{
				num1: "9",
				num2: "9",
			},
			expected: "18",
		},
		{
			input: &Input{
				num1: "1",
				num2: "9",
			},
			expected: "10",
		},
	}

	for _, tc := range testCases {
		actual := addStrings(tc.input.num1, tc.input.num2)
		if actual != tc.expected {
			fmt.Printf("addStrings(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
