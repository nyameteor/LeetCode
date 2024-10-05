package main

import "fmt"

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	s1Count := make([]int, 26)
	s2Count := make([]int, 26)

	for i := 0; i < len(s1); i++ {
		s1Count[s1[i]-'a']++
		s2Count[s2[i]-'a']++
	}

	for i := len(s1); i < len(s2); i++ {
		if matches(s1Count, s2Count) {
			return true
		}
		s2Count[s2[i]-'a']++
		s2Count[s2[i-len(s1)]-'a']--
	}

	return matches(s1Count, s2Count)
}

func matches(count1 []int, count2 []int) bool {
	for i := 0; i < 26; i++ {
		if count1[i] != count2[i] {
			return false
		}
	}
	return true
}

func main() {
	var s1 string
	var s2 string
	var res bool

	s1 = "ab"
	s2 = "eidbaooo"
	res = checkInclusion(s1, s2)
	// true
	fmt.Println(res)

	s1 = "ab"
	s2 = "eidboaoo"
	res = checkInclusion(s1, s2)
	// false
	fmt.Println(res)

	s1 = "a"
	s2 = "ab"
	res = checkInclusion(s1, s2)
	// true
	fmt.Println(res)

	s1 = "adc"
	s2 = "dcda"
	res = checkInclusion(s1, s2)
	// true
	fmt.Println(res)

	s1 = "ab"
	s2 = "a"
	res = checkInclusion(s1, s2)
	// false
	fmt.Println(res)
}
