package main

import "fmt"

// Recursion
// Time Limit Exceeded
func isMatchTle(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}

	if len(s) > 0 && (s[0] == p[0] || p[0] == '?') {
		return isMatch(s[1:], p[1:])
	}

	if p[0] == '*' {
		return (len(s) > 0 && isMatch(s[1:], p)) || isMatch(s, p[1:])
	}

	return false
}

// Memoization
// O(MN) space
func isMatch(s string, p string) bool {
	memo := map[int]map[int]bool{}

	lookup(s, p, 0, 0, memo)

	return memo[len(s)][len(p)]
}

func lookup(s string, p string, i int, j int, memo map[int]map[int]bool) bool {
	if _, ok := memo[i][j]; ok {
		return memo[i][j]
	}

	if j == len(p) {
		setMap(memo, i, j, i == len(s))
		return memo[i][j]
	}

	if i < len(s) && (s[i] == p[j] || p[j] == '?') {
		setMap(memo, i, j, lookup(s, p, i+1, j+1, memo))
		return memo[i][j]
	}

	if p[j] == '*' {
		setMap(memo, i, j, (i < len(s) && lookup(s, p, i+1, j, memo)) || lookup(s, p, i, j+1, memo))
		return memo[i][j]
	}

	setMap(memo, i, j, false)
	return memo[i][j]
}

// Help to set in nested map
func setMap(m map[int]map[int]bool, i int, j int, v bool) {
	if _, ok := m[i]; !ok {
		m[i] = make(map[int]bool)
	}
	m[i][j] = v
}

// Tabulation
// O(MN) space
func isMatch2(s string, p string) bool {
	n := len(s)
	m := len(p)
	table := make([][]bool, n+1)
	for i := 0; i < n+1; i++ {
		table[i] = make([]bool, m+1)
	}

	table[0][0] = true
	for j := 1; j < m+1 && p[j-1] == '*'; j++ {
		table[0][j] = table[0][j-1]
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			if p[j-1] == '*' {
				table[i][j] = table[i-1][j] || table[i][j-1]
			} else {
				table[i][j] = (s[i-1] == p[j-1] || p[j-1] == '?') && table[i-1][j-1]
			}
		}
	}

	return table[n][m]
}

func main() {
	var s string
	var p string

	s = "aabcbc"
	p = "aa*bc"
	fmt.Println(isMatch(s, p) == true)

	s = "acdcb"
	p = "a*c?b"
	fmt.Println(isMatch(s, p) == false)

	s = ""
	p = "*****"
	fmt.Println(isMatch(s, p) == true)

	s = "bbbbbbbabbaabbabbbbaaabbabbabaaabbababbbabbbabaaabaab"
	p = "b*b*ab**ba*b**b***bba"
	fmt.Println(isMatch(s, p) == false)
}
