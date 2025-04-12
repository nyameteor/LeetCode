package main

import (
	"fmt"
	"math/rand"
	"time"
)

var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

// Generates a uniform random integer in the range [1, 7]
func rand7() int {
	return seededRand.Intn(7) + 1
}

func rand10() int {
	a := rand7()
	b := rand7()
	index := (a-1)*7 + (b - 1)
	if index < 40 {
		return index%10 + 1
	}
	return rand10()
}

func main() {
	testCases := []struct {
		n int
	}{
		{n: 100},
		{n: 1000},
		{n: 10000},
	}

	for _, tc := range testCases {
		frequency := make([]int, 10)
		for range tc.n {
			res := rand10()
			frequency[res-1]++
		}
		fmt.Printf("n = %v\n", tc.n)
		for i, count := range frequency {
			fmt.Printf("Number %d appeared %d times (%.2f%%)\n", i+1, count, float64(count)*100/float64(tc.n))
		}
		fmt.Println()
	}
}
