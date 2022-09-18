package main

import "fmt"

func isInterleave(s1 string, s2 string, s3 string) bool {

	// edge case
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	memo := make([][]int, len(s1))
	for i := 0; i < len(s1); i++ {
		memo[i] = make([]int, len(s2))
	}

	return dfsIsInterleave(&s1, 0, &s2, 0, &s3, 0, &memo)
}

func dfsIsInterleave(s1 *string, i int, s2 *string, j int, s3 *string, k int, memo *[][]int) bool {
	if i == len(*s1) {
		return (*s2)[j:] == (*s3)[k:]
	}

	if j == len(*s2) {
		return (*s1)[i:] == (*s3)[k:]
	}

	if k == len(*s3) {
		return false
	}

	//  0: not memorized
	if (*memo)[i][j] != 0 {
		return getMemo(memo, i, j)
	}

	if (*s1)[i] == (*s2)[j] && (*s1)[i] == (*s3)[k] {
		ans := dfsIsInterleave(s1, i+1, s2, j, s3, k+1, memo) || dfsIsInterleave(s1, i, s2, j+1, s3, k+1, memo)
		setMemo(memo, i, j, ans)
		return ans
	}

	if (*s1)[i] == (*s3)[k] {
		ans := dfsIsInterleave(s1, i+1, s2, j, s3, k+1, memo)
		setMemo(memo, i, j, ans)
		return ans
	}

	if (*s2)[j] == (*s3)[k] {
		ans := dfsIsInterleave(s1, i, s2, j+1, s3, k+1, memo)
		setMemo(memo, i, j, ans)
		return ans
	}

	return false
}

// -1: false
// 	1: true
func getMemo(memo *[][]int, i int, j int) bool {
	if (*memo)[i][j] == -1 {
		return false
	} else {
		return true
	}
}

func setMemo(memo *[][]int, i int, j int, ans bool) {
	if ans {
		(*memo)[i][j] = 1
	} else {
		(*memo)[i][j] = -1
	}
}

func main() {
	var s1, s2, s3 string

	s1 = "aabcc"
	s2 = "dbbca"
	s3 = "aadbbcbcac"
	fmt.Println(isInterleave(s1, s2, s3))

	s1 = "aabcc"
	s2 = "dbbca"
	s3 = "aadbbbaccc"
	fmt.Println(isInterleave(s1, s2, s3))

	s1 = ""
	s2 = ""
	s3 = ""
	fmt.Println(isInterleave(s1, s2, s3))

	s1 = "aaaaaaaaaaaaaaaaaaaaaaaaaaa"
	s2 = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	s3 = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	fmt.Println(isInterleave(s1, s2, s3))
}
