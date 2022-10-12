package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	return kSum(nums, target, 4)
}

func kSum(nums []int, target int, k int) [][]int {
	res := [][]int{}

	if len(nums) == 0 {
		return res
	}

	if nums[0] > target/k || target/k > nums[len(nums)-1] {
		return res
	}

	if k == 2 {
		return twoSum(nums, target)
	}

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		for _, subset := range kSum(nums[i+1:], target-nums[i], k-1) {
			res = append(res, append([]int{nums[i]}, subset...))
		}
	}

	return res
}

func twoSum(nums []int, target int) [][]int {
	res := [][]int{}

	l := 0
	r := len(nums) - 1
	for l < r {
		if nums[l]+nums[r] < target || (l > 0 && nums[l] == nums[l-1]) {
			l++
		} else if nums[l]+nums[r] > target || (r < len(nums)-1 && nums[r] == nums[r+1]) {
			r--
		} else {
			res = append(res, []int{nums[l], nums[r]})
			l++
			r--
		}
	}
	return res
}

func main() {
	var nums []int
	var target int

	nums = []int{1, 0, -1, 0, -2, 2}
	target = 0
	fmt.Println(fourSum(nums, target))

	nums = []int{2, 2, 2, 2, 2}
	target = 8
	fmt.Println(fourSum(nums, target))
}
