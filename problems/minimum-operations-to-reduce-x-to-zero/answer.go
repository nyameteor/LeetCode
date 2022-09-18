package main

import (
	"fmt"
	"math"
)

// Sliding Window
func minOperations(nums []int, x int) int {
	n := len(nums)
	target := -x
	for i := 0; i < n; i++ {
		target += nums[i]
	}
	if target < 0 {
		return -1
	}

	sum := 0
	res := math.MinInt32
	l := 0
	r := 0
	for l < n && r < n {
		sum += nums[r]
		r++
		for sum > target {
			sum -= nums[l]
			l++
		}
		if sum == target {
			res = max(res, r-l)
		}
	}

	if res == math.MinInt32 {
		return -1
	} else {
		return n - res
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// Depth-First Search - Time Limit Exceeded
func minOperations1(nums []int, x int) int {
	n := len(nums)
	l := 0
	r := n - 1
	res := -1
	dfs(&nums, &res, 0, x, l, r)

	return res
}

func dfs(nums *[]int, res *int, step int, x int, l int, r int) {
	if x < 0 {
		return
	}

	if x == 0 {
		if *res == -1 {
			*res = step
		}
		*res = min(*res, step)
		return
	}

	if l > r {
		return
	}

	dfs(nums, res, step+1, x-(*nums)[l], l+1, r)
	dfs(nums, res, step+1, x-(*nums)[r], l, r-1)
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {
	nums := []int{1, 1, 4, 2, 3}
	x := 5
	fmt.Println(minOperations(nums, x))

	nums = []int{5, 6, 7, 8, 9}
	x = 4
	fmt.Println(minOperations(nums, x))

	nums = []int{3, 2, 20, 1, 1, 3}
	x = 10
	fmt.Println(minOperations(nums, x))

	nums = []int{8828, 9581, 49, 9818, 9974, 9869, 9991, 10000, 10000, 10000, 9999, 9993, 9904, 8819, 1231, 6309}
	x = 134365
	fmt.Println(minOperations(nums, x))
}
