package main

import "fmt"

func hasAllCodes(s string, k int) bool {
	n := len(s)
	if n < k {
		return false
	}

	need := 1 << k

	got := make(map[string]struct{})
	for i := 0; i <= n-k; i++ {
		sub := s[i : i+k]
		if _, ok := got[sub]; ok {
			continue
		} else {
			got[sub] = struct{}{}
			need--
			if need == 0 {
				return true
			}
		}
	}

	return false
}

func main() {
	s := "00110110"
	k := 2
	fmt.Println(hasAllCodes(s, k))

	s = "0110"
	k = 2
	fmt.Println(hasAllCodes(s, k))
}
