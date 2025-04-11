package main

import (
	"fmt"
)

func compareVersion(version1 string, version2 string) int {
	i, j := 0, 0

	for i < len(version1) || j < len(version2) {
		num1 := 0
		for i < len(version1) && version1[i] != '.' {
			num1 = num1*10 + int(version1[i]-'0')
			i++
		}
		i++

		num2 := 0
		for j < len(version2) && version2[j] != '.' {
			num2 = num2*10 + int(version2[j]-'0')
			j++
		}
		j++

		if num1 < num2 {
			return -1
		}
		if num1 > num2 {
			return 1
		}
	}

	return 0
}

func main() {
	type Input struct {
		version1, version2 string
	}

	testCases := []struct {
		input    *Input
		expected int
	}{
		{
			input: &Input{
				version1: "1.2",
				version2: "1.10",
			},
			expected: -1,
		},
		{
			input: &Input{
				version1: "1.01",
				version2: "1.001",
			},
			expected: 0,
		},
		{
			input: &Input{
				version1: "1.0",
				version2: "1.0.0.0",
			},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		actual := compareVersion(tc.input.version1, tc.input.version2)
		if actual != tc.expected {
			fmt.Printf("compareVersion(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
