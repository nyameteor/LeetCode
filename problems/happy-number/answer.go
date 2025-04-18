package main

import "fmt"

func isHappy(n int) bool {
	return isHappyFloyd(n)
}

func isHappySet(n int) bool {
	seen := map[int]bool{}
	for {
		if n == 1 {
			return true
		}
		if seen[n] {
			return false
		}
		seen[n] = true
		n = next(n)
	}
}

func isHappyFloyd(n int) bool {
	slow, fast := n, n
	for {
		slow = next(slow)
		fast = next(next(fast))
		if fast == 1 {
			return true
		}
		if slow == fast {
			return false
		}
	}
}

func next(n int) int {
	res := 0
	for n > 0 {
		d := n % 10
		res += d * d
		n /= 10
	}
	return res
}

func main() {
	testCases := []struct {
		input    int
		expected bool
	}{
		{
			input:    19,
			expected: true,
		},
		{
			input:    2,
			expected: false,
		},
		{
			input:    10,
			expected: true,
		},
	}

	solutions := []struct {
		name string
		fn   func(n int) bool
	}{
		{
			name: "Default Solution",
			fn:   isHappy,
		},
		{
			name: "Hash Set Method",
			fn:   isHappySet,
		},
		{
			name: "Floyd's Cycle Detection",
			fn:   isHappyFloyd,
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
