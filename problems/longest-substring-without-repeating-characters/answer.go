package main

import "fmt"

// Sliding Window
func lengthOfLongestSubstring(s string) int {
	maxLen := 0

	seen := make(map[byte]struct{})

	l := 0
	r := 0
	n := len(s)
	for l < n && r < n {
		if _, ok := seen[s[r]]; !ok {
			seen[s[r]] = struct{}{}
			r++
			maxLen = max(maxLen, r-l)
		} else {
			delete(seen, s[l])
			l++
		}
	}

	return maxLen
}

// Sliding Window (with Index)
func lengthOfLongestSubstring1(s string) int {
	maxLen := 0

	seen := map[byte]int{}

	l := 0
	r := 0
	for r < len(s) {
		if _, ok := seen[s[r]]; ok {
			l = max(l, seen[s[r]]+1)
		}
		seen[s[r]] = r
		maxLen = max(maxLen, r-l+1)
		r++
	}

	return maxLen
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))

	s = "pwwkew"
	fmt.Println(lengthOfLongestSubstring(s))

	s = "dvdf"
	fmt.Println(lengthOfLongestSubstring(s))

	s = "abba"
	fmt.Println(lengthOfLongestSubstring(s))

	s = "b"
	fmt.Println(lengthOfLongestSubstring(s))
}
