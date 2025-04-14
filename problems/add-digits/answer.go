package main

import "fmt"

func addDigits(num int) int {
	return digitalRootFormula(num)
}

func digitalRootRecursive(num int) int {
	if num < 10 {
		return num
	}
	return digitalRootRecursive(num/10 + num%10)
}

func digitalRootFormula(num int) int {
	if num == 0 {
		return 0
	}
	return 1 + (num-1)%9
}

func main() {
	testCases := []struct {
		input    int
		expected int
	}{
		{
			input:    38,
			expected: 2,
		},
		{
			input:    0,
			expected: 0,
		},
	}

	solutions := []struct {
		name string
		fn   func(int) int
	}{
		{
			name: "Default Solution",
			fn:   addDigits,
		},
		{
			name: "Recursive Simulation",
			fn:   digitalRootRecursive,
		},
		{
			name: "Mathematical Formula",
			fn:   digitalRootFormula,
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
