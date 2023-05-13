package main

import "fmt"

func maxUncrossedLines(nums1 []int, nums2 []int) int {
	var m = len(nums1)
	var n = len(nums2)

	const NOT_PRESENT = -1
	var memo = make([][]int, m)
	for i := 0; i < m; i++ {
		memo[i] = make([]int, n)
		for j := 0; j < n; j++ {
			memo[i][j] = NOT_PRESENT
		}
	}

	var lookup func(i, j int) int
	lookup = func(i, j int) int {
		if i < 0 || j < 0 {
			return 0
		}
		if memo[i][j] == NOT_PRESENT {
			if nums1[i] == nums2[j] {
				memo[i][j] = 1 + lookup(i-1, j-1)
			} else {
				memo[i][j] = max(lookup(i-1, j), lookup(i, j-1))
			}
		}
		return memo[i][j]
	}
	return lookup(m-1, n-1)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func main() {
	var nums1 []int
	var nums2 []int

	nums1 = []int{1, 4, 2}
	nums2 = []int{1, 2, 4}
	// 2
	fmt.Println(maxUncrossedLines(nums1, nums2))

	nums1 = []int{2, 5, 1, 2, 5}
	nums2 = []int{10, 5, 2, 1, 5, 2}
	// 3
	fmt.Println(maxUncrossedLines(nums1, nums2))

	nums1 = []int{1, 3, 7, 1, 7, 5}
	nums2 = []int{1, 9, 2, 5, 1}
	// 2
	fmt.Println(maxUncrossedLines(nums1, nums2))
}
