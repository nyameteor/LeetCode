package main

import (
	"fmt"
	"math"
)

// Recursion
// Time Limit Exceeded
func minDifficultyTle(jobDifficulty []int, d int) int {
	if len(jobDifficulty) < d {
		return -1
	}

	return lookupTle(jobDifficulty, d)
}

func lookupTle(jobDifficulty []int, d int) int {
	if d == 1 {
		cur := 0
		for i := 0; i < len(jobDifficulty); i++ {
			cur = max(cur, jobDifficulty[i])
		}
		return cur
	}

	cur := 0
	res := math.MaxInt32
	for i := 0; i < len(jobDifficulty)-d+1; i++ {
		cur = max(cur, jobDifficulty[i])
		res = min(res, cur+lookupTle(jobDifficulty[i+1:], d-1))
	}
	return res
}

// Memoization
func minDifficulty(jobDifficulty []int, d int) int {
	if len(jobDifficulty) < d {
		return -1
	}

	memo := map[int]map[int]int{}
	return lookup(jobDifficulty, d, 0, 0, memo)
}

func lookup(jobDifficulty []int, d int, i int, j int, memo map[int]map[int]int) int {
	if _, ok := memo[i][j]; ok {
		return memo[i][j]
	}

	if j == d-1 {
		cur := 0
		for k := i; k < len(jobDifficulty); k++ {
			cur = max(cur, jobDifficulty[k])
		}
		setMap(memo, i, j, cur)
		return memo[i][j]
	}

	cur := 0
	res := math.MaxInt32
	for k := i; k < len(jobDifficulty)-(d-j-1); k++ {
		cur = max(cur, jobDifficulty[k])
		res = min(res, cur+lookup(jobDifficulty, d, k+1, j+1, memo))
	}
	setMap(memo, i, j, res)
	return memo[i][j]
}

// Help to set in nested map
func setMap(m map[int]map[int]int, i int, j int, v int) {
	if _, ok := m[i]; !ok {
		m[i] = make(map[int]int)
	}
	m[i][j] = v
}

// Helper
func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	var jobDifficulty []int
	var d int

	jobDifficulty = []int{6, 5, 4, 3, 2, 1}
	d = 2
	fmt.Println(minDifficulty(jobDifficulty, d))

	jobDifficulty = []int{6, 5, 4, 3, 2, 1}
	d = 3
	fmt.Println(minDifficulty(jobDifficulty, d))

	jobDifficulty = []int{7, 1, 7, 1, 7, 1}
	d = 3
	fmt.Println(minDifficulty(jobDifficulty, d))

	jobDifficulty = []int{9, 9, 9}
	d = 4
	fmt.Println(minDifficulty(jobDifficulty, d))

	jobDifficulty = []int{380, 302, 102, 681, 863, 676, 243, 671, 651,
		612, 162, 561, 394, 856, 601, 30, 6, 257, 921, 405, 716, 126,
		158, 476, 889, 699, 668, 930, 139, 164, 641, 801, 480, 756, 797,
		915, 275, 709, 161, 358, 461, 938, 914, 557, 121, 964, 315}
	d = 10
	fmt.Println(minDifficulty(jobDifficulty, d))
}
