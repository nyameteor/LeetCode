package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	l := 0
	r := 0
	words := []string{}
	for r < len(s) {
		// let `s[l]` be the first letter of current word
		for l < len(s) && s[l] == ' ' {
			l++
		}
		// corner case: `l` reach the end of string
		if l == len(s) {
			break
		}
		r = l + 1

		// let `s[r]` be the first space after current word
		for r < len(s) && s[r] != ' ' {
			r++
		}
		words = append(words, s[l:r])
		l = r + 1
	}

	reverse(words)
	return strings.Join(words, " ")
}

func reverse(words []string) {
	l := 0
	r := len(words) - 1
	for l < r {
		words[l], words[r] = words[r], words[l]
		l++
		r--
	}
}

func main() {
	var s string

	s = "the sky is blue"
	fmt.Println(reverseWords(s))

	s = "  hello world  "
	fmt.Println(reverseWords(s))

	s = "a good   example"
	fmt.Println(reverseWords(s))

	s = "F R  I   E    N     D      S      "
	fmt.Println(reverseWords(s))
}
