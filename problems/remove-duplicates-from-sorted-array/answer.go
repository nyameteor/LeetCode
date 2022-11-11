package main

import "fmt"

func removeDuplicates(nums []int) int {
	l := 0
	r := 1
	for r < len(nums) {
		if nums[l] != nums[r] {
			l++
			nums[l] = nums[r]
		}
		r++
	}
	return l + 1
}

func main() {
	var nums []int

	nums = []int{1, 1, 2}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)

	nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)
}
