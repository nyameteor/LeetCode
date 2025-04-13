package main

import "fmt"

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	res := 0
	sumFreq := make(map[int]int)

	for _, a := range nums1 {
		for _, b := range nums2 {
			sumFreq[a+b]++
		}
	}

	for _, c := range nums3 {
		for _, d := range nums4 {
			res += sumFreq[-(c + d)]
		}
	}

	return res
}

func main() {
	type Input struct {
		nums1, nums2, nums3, nums4 []int
	}

	testCases := []struct {
		input    *Input
		expected int
	}{
		{
			input: &Input{
				nums1: []int{1, 2},
				nums2: []int{-2, -1},
				nums3: []int{-1, 2},
				nums4: []int{0, 2},
			},
			expected: 2,
		},
		{
			input: &Input{
				nums1: []int{0},
				nums2: []int{0},
				nums3: []int{0},
				nums4: []int{0},
			},
			expected: 1,
		},
	}

	for _, tc := range testCases {
		actual := fourSumCount(tc.input.nums1, tc.input.nums2, tc.input.nums3, tc.input.nums4)
		if actual != tc.expected {
			fmt.Printf("fourSumCount(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
