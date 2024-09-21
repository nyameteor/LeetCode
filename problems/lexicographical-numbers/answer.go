package main

import (
	"fmt"
)

func lexicalOrder(n int) []int {
	res := []int{}
	helper(0, 0, n, &res)
	return res
}

func helper(pre int, depth int, n int, res *[]int) {
	for i := 0; i <= 9; i++ {
		if depth == 0 && i == 0 {
			continue
		}
		cur := pre*10 + i
		if cur <= n {
			*res = append(*res, cur)
			helper(cur, depth+1, n, res)
		}
	}
}

func main() {
	var n int
	var res []int

	n = 13
	res = lexicalOrder(n)
	// [1,10,11,12,13,2,3,4,5,6,7,8,9]
	fmt.Println(res)

	n = 2
	res = lexicalOrder(n)
	// [1,2]
	fmt.Println(res)
}
