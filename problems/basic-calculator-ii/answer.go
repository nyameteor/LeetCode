package main

import (
	"fmt"
)

func calculate(s string) int {
	ops := []byte{}
	nums := []int{}

	for i := 0; i < len(s); {
		if isSpace(s[i]) {
			i++
		} else if isOperator(s[i]) {
			for len(ops) > 0 && precedence(s[i]) <= precedence(ops[len(ops)-1]) {
				applyTopOperator(&ops, &nums)
			}
			ops = append(ops, s[i])
			i++
		} else {
			num := 0
			for i < len(s) && isDigit(s[i]) {
				num = num*10 + int(s[i]-'0')
				i++
			}
			nums = append(nums, num)
		}
	}

	for len(ops) > 0 {
		applyTopOperator(&ops, &nums)
	}

	return nums[0]
}

func applyTopOperator(ops *[]byte, nums *[]int) {
	op := (*ops)[len(*ops)-1]
	*ops = (*ops)[:len(*ops)-1]

	right, left := (*nums)[len(*nums)-1], (*nums)[len(*nums)-2]
	*nums = (*nums)[:len(*nums)-2]

	result := evaluate(op, left, right)
	*nums = append(*nums, result)
}

func evaluate(op byte, left, right int) int {
	switch op {
	case '+':
		return left + right
	case '-':
		return left - right
	case '*':
		return left * right
	case '/':
		return left / right
	default:
		panic("Invalid operator")
	}
}

func precedence(op byte) int {
	switch op {
	case '+', '-':
		return 1
	case '*', '/':
		return 2
	default:
		panic("Invalid operator")
	}
}

func isSpace(ch byte) bool {
	return ch == ' '
}

func isOperator(ch byte) bool {
	return ch == '+' || ch == '-' || ch == '*' || ch == '/'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func main() {
	testCases := []struct {
		input    string
		expected int
	}{
		{
			input:    "3+2*2",
			expected: 7,
		},
		{
			input:    " 3/2 ",
			expected: 1,
		},
		{
			input:    " 3+5 / 2 ",
			expected: 5,
		},
	}

	for _, tc := range testCases {
		actual := calculate(tc.input)
		if actual != tc.expected {
			fmt.Printf("calculate(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
