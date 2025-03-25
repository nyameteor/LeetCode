package main

import "fmt"

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		m := l + (r-l)/2
		if nums[m] == target {
			return m
		}

		if nums[m] >= nums[l] {
			// Left part is sorted, or pivot is in the right part
			if nums[m] > target && nums[l] <= target {
				// Target is in the left part, discard right part
				r = m - 1
			} else {
				// Otherwise, discard the left part
				l = m + 1
			}
		} else {
			// Right part is sorted, or pivot is in the left part
			if nums[m] < target && nums[r] >= target {
				// Target is in the right part, discard left part
				l = m + 1
			} else {
				// Otherwise, discard the right part
				r = m - 1
			}
		}
	}

	return -1
}

func main() {
	type Input struct {
		nums   []int
		target int
	}

	testCases := []struct {
		input    *Input
		expected int
	}{
		{
			input: &Input{
				nums:   []int{4, 5, 6, 7, 0, 1, 2},
				target: 0,
			},
			expected: 4,
		},
		{
			input: &Input{
				nums:   []int{4, 5, 6, 7, 0, 1, 2},
				target: 3,
			},
			expected: -1,
		},
		{
			input: &Input{
				nums:   []int{1},
				target: 0,
			},
			expected: -1,
		},
		{
			input: &Input{
				nums:   []int{6, 7, 0, 1, 2},
				target: 7,
			},
			expected: 1,
		},
		{
			input: &Input{
				nums:   []int{3, 1},
				target: 1,
			},
			expected: 1,
		},
	}

	for _, tc := range testCases {
		actual := search(tc.input.nums, tc.input.target)
		if actual != tc.expected {
			fmt.Printf("search(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
