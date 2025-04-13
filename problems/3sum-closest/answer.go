package main

import (
	"fmt"
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	res := 0
	minDist := math.MaxInt32

	for i := 0; i < len(nums)-2; i++ {
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			dist := abs(sum - target)

			if dist < minDist {
				minDist = dist
				res = sum
			}

			if sum == target {
				break
			} else if sum < target {
				l++
			} else {
				r--
			}
		}
	}

	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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
				nums:   []int{-1, 2, 1, -4},
				target: 1,
			},
			expected: 2,
		},
		{
			input: &Input{
				nums:   []int{0, 0, 0},
				target: 1,
			},
			expected: 0,
		},
		{
			input: &Input{
				nums:   []int{0, 2, 1, -3},
				target: 1,
			},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		actual := threeSumClosest(tc.input.nums, tc.input.target)
		if actual != tc.expected {
			fmt.Printf("threeSumClosest(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
