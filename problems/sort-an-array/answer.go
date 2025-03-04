package main

import (
	"fmt"
	"math"
	"slices"
)

// Selection sort: Find the smallest item and put it at the front.
func selectionSort(nums []int) {
	n := len(nums)

	for i := 0; i < n; i++ {
		minNum := math.MaxInt32
		minIdx := -1
		for j := i; j < n; j++ {
			if nums[j] < minNum {
				minNum = nums[j]
				minIdx = j
			}
		}
		nums[i], nums[minIdx] = nums[minIdx], nums[i]
	}
}

// Insertion sort: Insert current item into its correct position in the sorted portion.
func insertionSort(nums []int) {
	n := len(nums)

	for i := 0; i < n; i++ {
		k := i
		for j := i - 1; j >= 0; j-- {
			if nums[k] < nums[j] {
				nums[k], nums[j] = nums[j], nums[k]
				k = j
			} else {
				break
			}
		}
	}
}

// Merge Sort: Merge two sorted halves into one sorted whole.
func mergeSort(nums []int) {
	merge := func(left, right []int) []int {
		m := len(left)
		n := len(right)

		result := make([]int, m+n)
		i, j, k := 0, 0, 0
		for i < m && j < n {
			if left[i] < right[j] {
				result[k] = left[i]
				k++
				i++
			} else {
				result[k] = right[j]
				k++
				j++
			}
		}

		for i < m {
			result[k] = left[i]
			i++
			k++
		}
		for j < n {
			result[k] = right[j]
			k++
			j++
		}

		return result
	}

	var mergeSortRecursive func(arr []int) []int
	mergeSortRecursive = func(arr []int) []int {
		n := len(arr)
		if n <= 1 {
			return arr
		}

		mid := n / 2
		left := mergeSortRecursive(arr[:mid])
		right := mergeSortRecursive(arr[mid:])
		return merge(left, right)
	}

	result := mergeSortRecursive(nums)
	copy(nums, result)
}

// Partition(quick) sort: Partition items around a pivot.
func quickSort(nums []int) {
	partition := func(arr []int, left, right int) int {
		pivot := arr[left]

		i := left - 1
		j := right + 1

		for {
			for i++; arr[i] < pivot; i++ {
			}
			for j--; arr[j] > pivot; j-- {
			}
			if i >= j {
				return j
			}
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	var quickSortRecursive func(arr []int, left, right int)
	quickSortRecursive = func(arr []int, left, right int) {
		if left >= right {
			return
		}

		mid := partition(arr, left, right)
		quickSortRecursive(arr, left, mid)
		quickSortRecursive(arr, mid+1, right)
	}

	quickSortRecursive(nums, 0, len(nums)-1)
}

func main() {
	testCases := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{5, 2, 3, 1},
			expected: []int{1, 2, 3, 5},
		},
		{
			input:    []int{5, 1, 1, 2, 0, 0},
			expected: []int{0, 0, 1, 1, 2, 5},
		},
		{
			input:    []int{3, 7, 8, 5, 2, 1, 9, 5, 4},
			expected: []int{1, 2, 3, 4, 5, 5, 7, 8, 9},
		},
		{
			input:    []int{6, 5, 3, 1, 8, 7, 2, 4},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			input:    []int{3, 7, 4, 9, 5, 2, 6, 1},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 9},
		},
	}

	type SortFunc func(nums []int)

	sortSolutions := []struct {
		name string
		fn   SortFunc
	}{
		{
			name: "Selection Sort",
			fn:   selectionSort,
		},
		{
			name: "Insertion Sort",
			fn:   insertionSort,
		},
		{
			name: "Merge Sort",
			fn:   mergeSort,
		},
		{
			name: "Quick Sort (Partition Sort)",
			fn:   quickSort,
		},
	}

	for _, solution := range sortSolutions {
		for _, tc := range testCases {
			inputCopy := slices.Clone(tc.input)
			solution.fn(tc.input)
			if !slices.Equal(tc.input, tc.expected) {
				fmt.Printf("[FAILED] Name: %s | Input: %v | Output: %v | Expected: %v\n",
					solution.name, inputCopy, tc.input, tc.expected)
			}
			copy(tc.input, inputCopy)
		}
	}
}
