package main

import (
	"fmt"
	"reflect"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}

	for i := 0; i < len(nums)-2; i++ {
		// skip duplicates
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				l++
				r--

				// skip duplicates
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			} else if sum < 0 {
				l++
			} else {
				r--
			}
		}
	}

	return res
}

func main() {
	testCases := []struct {
		input    []int
		expected [][]int
	}{
		{
			input:    []int{-1, 0, 1, 2, -1, -4},
			expected: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			input:    []int{0, 1, 1},
			expected: [][]int{},
		},
		{
			input:    []int{0, 0, 0},
			expected: [][]int{{0, 0, 0}},
		},
	}

	for _, tc := range testCases {
		actual := threeSum(tc.input)
		if !reflect.DeepEqual(actual, tc.expected) {
			fmt.Printf("threeSum(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
