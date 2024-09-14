package main

import "fmt"

func longestSubarray(nums []int) int {
	count, maxCount := 1, 1
	maxNum := 0
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
			count, maxCount = 1, 1
		} else if num < maxNum {
			count = 0
		} else {
			count++
			if count > maxCount {
				maxCount = count
			}
		}
	}

	return maxCount
}

func main() {
	var nums []int
	var res int

	nums = []int{1, 2, 3, 3, 2, 2}
	res = longestSubarray(nums)
	// 2
	fmt.Println(res)

	nums = []int{1, 2, 3, 4}
	res = longestSubarray(nums)
	// 1
	fmt.Println(res)

	nums = []int{311155, 311155, 311155, 311155, 311155, 311155, 311155, 311155, 201191, 311155}
	res = longestSubarray(nums)
	// 8
	fmt.Println(res)

	nums = []int{96317, 96317, 96317, 96317, 96317, 96317, 96317, 96317, 96317, 279979}
	res = longestSubarray(nums)
	// 1
	fmt.Println(res)
}
