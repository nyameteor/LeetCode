package main

import "fmt"

func canArrange(arr []int, k int) bool {
	freq := make([]int, k)
	for _, e := range arr {
		r := e % k
		if r < 0 {
			r += k
		}
		freq[r]++
	}

	if freq[0]%2 != 0 {
		return false
	}

	for i := 1; i <= k-1; i++ {
		if freq[i] != freq[k-i] {
			return false
		}
	}

	return true
}

func main() {
	var arr []int
	var k int
	var res bool

	arr = []int{1, 2, 3, 4, 5, 10, 6, 7, 8, 9}
	k = 5
	res = canArrange(arr, k)
	// true
	fmt.Println(res)

	arr = []int{1, 2, 3, 4, 5, 6}
	k = 7
	res = canArrange(arr, k)
	// true
	fmt.Println(res)

	arr = []int{1, 2, 3, 4, 5, 6}
	k = 10
	res = canArrange(arr, k)
	// false
	fmt.Println(res)

	arr = []int{-1, 1, -2, 2, -3, 3, -4, 4}
	k = 3
	res = canArrange(arr, k)
	// true
	fmt.Println(res)
}
