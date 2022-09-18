package main

import "fmt"

func minDistance(word1 string, word2 string) int {
	len1 := len(word1)
	len2 := len(word2)
	memo := init2dArray(len1, len2)

	return len1 + len2 - 2*lcs(&memo, &word1, &word2, len1-1, len2-1)
}

// Depth-First Search - with Memoization
// Get longest common subsequence
func lcs(memo *[][]int, s1 *string, s2 *string, i1 int, i2 int) int {
	if i1 == -1 || i2 == -1 {
		return 0
	}

	if (*memo)[i1][i2] > 0 {
		return (*memo)[i1][i2]
	}

	if (*s1)[i1] == (*s2)[i2] {
		(*memo)[i1][i2] = 1 + lcs(memo, s1, s2, i1-1, i2-1)
	} else {
		(*memo)[i1][i2] = max(lcs(memo, s1, s2, i1-1, i2), lcs(memo, s1, s2, i1, i2-1))
	}
	return (*memo)[i1][i2]
}

func init2dArray(m int, n int) [][]int {
	arr := make([][]int, m)
	for i := 0; i < m; i++ {
		arr[i] = make([]int, n)
	}
	return arr
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	s1 := "sea"
	s2 := "eat"
	fmt.Println(minDistance(s1, s2))

	s1 = "leetcode"
	s2 = "etco"
	fmt.Println(minDistance(s1, s2))
}
