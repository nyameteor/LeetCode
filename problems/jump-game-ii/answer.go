package main

import (
	"fmt"
	"math"
)

// Greedy
func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	n := len(nums)
	coverPos := 0
	step := 0

	for curPos := 0; curPos <= coverPos && curPos < n; {
		if coverPos < curPos+nums[curPos] {
			coverPos = curPos + nums[curPos]
			step++
		}
		if coverPos >= n-1 {
			return step
		}

		nextPos := 0
		maxDistance := 0
		for j := curPos + 1; j <= coverPos; j++ {
			if j+nums[j] > maxDistance {
				maxDistance = j + nums[j]
				nextPos = j
			}
		}
		curPos = nextPos
	}

	return step
}

// Depth-First Search
func jump2(nums []int) int {
	n := len(nums)
	memo := map[int]int{}

	return dfs(&memo, &nums, &n, 0)
}

func dfs(memo *map[int]int, nums *[]int, n *int, x int) int {
	if val, ok := (*memo)[x]; ok {
		return val
	}

	if x >= *n-1 {
		return 0
	}

	var minJump = math.MaxInt32
	for i := 1; i <= (*nums)[x]; i++ {
		minJump = min(minJump, 1+dfs(memo, nums, n, x+i))
	}

	(*memo)[x] = minJump
	return minJump
}

func min(a int, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func main() {
	nums := []int{2, 3, 1, 1, 4}
	fmt.Println(jump(nums))

	nums = []int{2, 3, 0, 1, 4}
	fmt.Println(jump(nums))

	nums = []int{1, 2, 1, 1, 1}
	fmt.Println(jump(nums))

	nums = []int{7, 0, 9, 6, 9, 6, 1, 7, 9, 0, 1, 2, 9, 0, 3}
	fmt.Println(jump(nums))
}
