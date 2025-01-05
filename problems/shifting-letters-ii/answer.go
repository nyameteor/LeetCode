package main

import "fmt"

func shiftingLetters(s string, shifts [][]int) string {
	n := len(s)
	diffArr := make([]int, n+1)

	for _, shift := range shifts {
		start, end, direction := shift[0], shift[1], shift[2]
		if direction == 0 {
			diffArr[start] -= 1
			diffArr[end+1] += 1
		} else {
			diffArr[start] += 1
			diffArr[end+1] -= 1
		}
	}

	result := []byte(s)
	curShift := 0
	for i := 0; i < n; i++ {
		curShift += diffArr[i]
		result[i] = 'a' + byte((int(result[i]-'a')+curShift)%26+26)%26
	}

	res := string(result)
	return res
}

func main() {
	type Input struct {
		s      string
		shifts [][]int
	}
	testCases := []struct {
		input    Input
		expected string
	}{
		{
			input: Input{
				s: "abc",
				shifts: [][]int{
					{0, 1, 0}, {1, 2, 1}, {0, 2, 1},
				},
			},
			expected: "ace",
		},
		{
			input: Input{
				s: "dztz",
				shifts: [][]int{
					{0, 0, 0}, {1, 1, 1},
				},
			},
			expected: "catz",
		},
		{
			input: Input{
				s: "xuwdbdqik",
				shifts: [][]int{
					{4, 8, 0}, {4, 4, 0}, {2, 4, 0}, {2, 4, 0}, {6, 7, 1}, {2, 2, 1}, {0, 2, 1}, {8, 8, 0}, {1, 3, 1},
				},
			},
			expected: "ywxcxcqii",
		},
	}

	for _, tc := range testCases {
		actual := shiftingLetters(tc.input.s, tc.input.shifts)
		if actual != tc.expected {
			fmt.Printf("shiftingLetters(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
