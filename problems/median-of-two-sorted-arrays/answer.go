package main

import (
	"fmt"
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	return binarySearchIterative(nums1, nums2)
}

func binarySearchRecursive(nums1 []int, nums2 []int) float64 {
	// Find the k-th smallest element in the union of two sorted arrays
	var findKth func(a, b []int, k int) int
	findKth = func(a, b []int, k int) int {
		if len(a) > len(b) {
			return findKth(b, a, k)
		}

		if len(a) == 0 {
			return b[k-1]
		}

		if k == 1 {
			return min(a[0], b[0])
		}

		i := min(k/2, len(a))
		j := k - i

		if a[i-1] < b[j-1] {
			return findKth(a[i:], b, k-i)
		}
		return findKth(a, b[j:], k-j)
	}

	total := len(nums1) + len(nums2)
	if total%2 == 1 {
		return float64(findKth(nums1, nums2, total/2+1))
	}
	left := findKth(nums1, nums2, total/2)
	right := findKth(nums1, nums2, total/2+1)
	return float64(left+right) / 2
}

func binarySearchIterative(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)

	if m > n {
		return binarySearchIterative(nums2, nums1)
	}

	low, high := 0, m
	for low <= high {
		i := (low + high) / 2
		j := (m+n+1)/2 - i

		leftMax1 := math.MinInt32
		if i > 0 {
			leftMax1 = nums1[i-1]
		}

		leftMax2 := math.MinInt32
		if j > 0 {
			leftMax2 = nums2[j-1]
		}

		rightMin1 := math.MaxInt32
		if i < m {
			rightMin1 = nums1[i]
		}

		rightMin2 := math.MaxInt32
		if j < n {
			rightMin2 = nums2[j]
		}

		if leftMax1 > rightMin2 {
			high = i - 1
		} else if leftMax2 > rightMin1 {
			low = i + 1
		} else {
			if (m+n)%2 != 0 {
				return float64(max(leftMax1, leftMax2))
			} else {
				return float64(max(leftMax1, leftMax2)+min(rightMin1, rightMin2)) / 2
			}
		}
	}

	return -1
}

func main() {
	type Input struct {
		nums1, nums2 []int
	}

	testCases := []struct {
		input    *Input
		expected float64
	}{
		{
			input: &Input{
				nums1: []int{1, 3},
				nums2: []int{2},
			},
			expected: 2.00000,
		},
		{
			input: &Input{
				nums1: []int{1, 2},
				nums2: []int{3, 4},
			},
			expected: 2.50000,
		},
		{
			input: &Input{
				nums1: []int{},
				nums2: []int{1},
			},
			expected: 1.0000,
		},
	}

	solutions := []struct {
		name string
		fn   func(nums1, nums2 []int) float64
	}{
		{
			name: "Default Solution",
			fn:   findMedianSortedArrays,
		},
		{
			name: "Binary Search (Recursive)",
			fn:   binarySearchRecursive,
		},
		{
			name: "Binary Search (Iterative)",
			fn:   binarySearchIterative,
		},
	}

	floatEqual := func(a, b, epsilon float64) bool {
		return math.Abs(a-b) <= epsilon
	}
	epsilon := 1e-9

	for _, solution := range solutions {
		for _, tc := range testCases {
			actual := solution.fn(tc.input.nums1, tc.input.nums2)
			if !floatEqual(actual, tc.expected, epsilon) {
				fmt.Printf("[FAILED] Name: %s | Input: %v | Output: %v | Expected: %v\n",
					solution.name, tc.input, actual, tc.expected)
			}
		}
	}
}
