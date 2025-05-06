package main

import "fmt"

func trap(height []int) int {
	return bottomUpDp(height)
}

func bottomUpDp(height []int) int {
	n := len(height)

	maxLefts := make([]int, n)
	maxLefts[0] = height[0]
	for i := 1; i < n; i++ {
		maxLefts[i] = max(maxLefts[i-1], height[i])
	}

	maxRights := make([]int, n)
	maxRights[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		maxRights[i] = max(maxRights[i+1], height[i])
	}

	res := 0
	for i := range height {
		res += min(maxLefts[i], maxRights[i]) - height[i]
	}

	return res
}

func twoPointers(height []int) int {
	res := 0

	l, r := 0, len(height)-1
	maxLeft, maxRight := 0, 0

	for l <= r {
		if height[l] <= height[r] {
			maxLeft = max(maxLeft, height[l])
			res += maxLeft - height[l]
			l++
		} else {
			maxRight = max(maxRight, height[r])
			res += maxRight - height[r]
			r--
		}
	}

	return res
}

func main() {
	testCases := []struct {
		input    []int
		expected int
	}{
		{
			input:    []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			expected: 6,
		},
		{
			input:    []int{4, 2, 0, 3, 2, 5},
			expected: 9,
		},
	}

	solutions := []struct {
		name string
		fn   func([]int) int
	}{
		{
			name: "Default Solution",
			fn:   trap,
		},
		{
			name: "Bottom-up DP",
			fn:   bottomUpDp,
		},
		{
			name: "Two Pointers",
			fn:   twoPointers,
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
