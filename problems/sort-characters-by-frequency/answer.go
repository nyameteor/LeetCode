package main

import (
	"fmt"
	"sort"
)

func frequencySort(s string) string {
	return bucketSort(s)
}

func bucketSort(s string) string {
	freq := make(map[byte]int)
	for i := range s {
		freq[s[i]]++
	}

	buckets := make([][]byte, len(s)+1)
	for k, v := range freq {
		buckets[v] = append(buckets[v], k)
	}

	res := make([]byte, 0, len(s))
	for i := len(buckets) - 1; i >= 0; i-- {
		for _, k := range buckets[i] {
			for range i {
				res = append(res, k)
			}
		}
	}

	return string(res)
}

func basicSort(s string) string {
	freq := make(map[byte]int)
	for i := range s {
		freq[s[i]]++
	}

	keys := []byte{}
	for k := range freq {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return freq[keys[i]] > freq[keys[j]]
	})

	res := make([]byte, 0, len(s))
	for _, k := range keys {
		for range freq[k] {
			res = append(res, k)
		}
	}

	return string(res)
}

func main() {
	testCases := []struct {
		input    string
		expected map[string]struct{}
	}{
		{
			input:    "tree",
			expected: map[string]struct{}{"eert": {}, "eetr": {}},
		},
		{
			input:    "cccaaa",
			expected: map[string]struct{}{"aaaccc": {}, "cccaaa": {}},
		},
		{
			input:    "Aabb",
			expected: map[string]struct{}{"bbAa": {}, "bbaA": {}},
		},
	}

	solutions := []struct {
		name string
		fn   func(s string) string
	}{
		{
			name: "Default Solution",
			fn:   frequencySort,
		},
		{
			name: "Basic Sort",
			fn:   basicSort,
		},
		{
			name: "Bucket Sort",
			fn:   bucketSort,
		},
	}

	for _, solution := range solutions {
		for _, tc := range testCases {
			actual := solution.fn(tc.input)
			if _, ok := tc.expected[actual]; !ok {
				fmt.Printf("[FAILED] Name: %s | Input: %v | Output: %v | Expected: %v\n",
					solution.name, tc.input, actual, tc.expected)
			}
		}
	}
}
