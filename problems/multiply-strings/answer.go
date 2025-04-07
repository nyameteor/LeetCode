package main

import "fmt"

func multiply(num1 string, num2 string) string {
	return multiplyNumsOptimized(num1, num2)
}

func multiplyNumsOptimized(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	m, n := len(num1), len(num2)
	res := make([]byte, m+n)

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			res[i+j+1] += (num1[i] - '0') * (num2[j] - '0')
			if res[i+j+1] >= 10 {
				res[i+j] += res[i+j+1] / 10
				res[i+j+1] %= 10
			}
		}
	}

	for res[0] == 0 {
		res = res[1:]
	}

	for i := range res {
		res[i] += '0'
	}

	return string(res)
}

func multiplyNumsDigitByDigit(num1 string, num2 string) string {
	reverse := func(s []byte) {
		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			s[i], s[j] = s[j], s[i]
		}
	}

	multiplySingleDigit := func(num string, digit byte, shift int) string {
		res := []byte{}
		carry := 0
		y := int(digit - '0')

		for i := len(num) - 1; i >= 0; i-- {
			x := int(num[i] - '0')
			product := x*y + carry
			res = append(res, byte('0'+product%10))
			carry = product / 10
		}
		if carry > 0 {
			res = append(res, byte('0'+carry))
		}

		reverse(res)
		for range shift {
			res = append(res, '0')
		}
		return string(res)
	}

	addNums := func(num1 string, num2 string) string {
		res := []byte{}
		carry := 0
		i, j := len(num1)-1, len(num2)-1

		for i >= 0 || j >= 0 {
			x, y := 0, 0
			if i >= 0 {
				x += int(num1[i] - '0')
				i--
			}
			if j >= 0 {
				y += int(num2[j] - '0')
				j--
			}
			sum := x + y + carry
			res = append(res, byte('0'+sum%10))
			carry = sum / 10
		}
		if carry > 0 {
			res = append(res, byte('0'+carry))
		}

		reverse(res)
		return string(res)
	}

	res := "0"
	if num1 == "0" || num2 == "0" {
		return res
	}

	for i := len(num2) - 1; i >= 0; i-- {
		product := multiplySingleDigit(num1, num2[i], len(num2)-1-i)
		res = addNums(res, product)
	}

	return res
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
				num1: "2",
				num2: "3",
			},
			expected: "6",
		},
		{
			input: &Input{
				num1: "123",
				num2: "456",
			},
			expected: "56088",
		},
		{
			input: &Input{
				num1: "9133",
				num2: "0",
			},
			expected: "0",
		},
	}

	solutions := []struct {
		name string
		fn   func(num1 string, num2 string) string
	}{
		{
			name: "Default Solution",
			fn:   multiply,
		},
		{
			name: "Digit-by-Digit Multiplication",
			fn:   multiplyNumsDigitByDigit,
		},
		{
			name: "Optimized Multiplication",
			fn:   multiplyNumsOptimized,
		},
	}

	for _, solution := range solutions {
		for _, tc := range testCases {
			actual := solution.fn(tc.input.num1, tc.input.num2)
			if actual != tc.expected {
				fmt.Printf("[FAILED] Name: %s | Input: %v | Output: %v | Expected: %v\n",
					solution.name, tc.input, actual, tc.expected)
			}
		}
	}
}
