package main

import "fmt"

func subarraySum(nums []int, k int) int {
	res := 0
	sum := 0
	sumFreq := make(map[int]int)
	sumFreq[0] = 1

	for i := range nums {
		sum += nums[i]
		if count, ok := sumFreq[sum-k]; ok {
			res += count
		}
		sumFreq[sum]++
	}

	return res
}

func main() {
	type Input struct {
		nums []int
		k    int
	}

	testCases := []struct {
		input    *Input
		expected int
	}{
		{
			input: &Input{
				nums: []int{1, 1, 1},
				k:    2,
			},
			expected: 2,
		},
		{
			input: &Input{
				nums: []int{1, 2, 3},
				k:    3,
			},
			expected: 2,
		},
		{
			input: &Input{
				nums: []int{1, 2, 1, 2, 1},
				k:    3,
			},
			expected: 4,
		},
	}

	for _, tc := range testCases {
		actual := subarraySum(tc.input.nums, tc.input.k)
		if actual != tc.expected {
			fmt.Printf("subarraySum(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
