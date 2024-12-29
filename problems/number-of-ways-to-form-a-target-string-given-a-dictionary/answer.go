package main

import "fmt"

func numWays(words []string, target string) int {
	const mod = 1e9 + 7

	freqs := computeFreqByIndex(words)

	n := len(target)
	m := len(freqs)

	memo := initMemo(n, m)

	res := helper(target, n, 0, freqs, m, 0, memo, mod)

	return res
}

func helper(target string, n int, i int, freqs []map[byte]int, m int, k int, memo [][]int, mod int) int {
	if i >= n {
		return 1
	}

	if k >= m {
		return 0
	}

	if m-k < n-i {
		return 0
	}

	if memo[i][k] != -1 {
		return memo[i][k]
	}

	ways := 0
	if count, ok := freqs[k][target[i]]; ok {
		ways += count * helper(target, n, i+1, freqs, m, k+1, memo, mod)
		ways %= mod
	}

	ways += helper(target, n, i, freqs, m, k+1, memo, mod)
	ways %= mod

	memo[i][k] = ways
	return ways
}

func computeFreqByIndex(words []string) []map[byte]int {
	res := []map[byte]int{}

	n := len(words)
	m := len(words[0])

	for j := 0; j < m; j++ {
		freq := map[byte]int{}
		for i := 0; i < n; i++ {
			freq[words[i][j]]++
		}
		res = append(res, freq)
	}

	return res
}

func initMemo(n int, m int) [][]int {
	memo := make([][]int, n)

	for i := 0; i < n; i++ {
		memo[i] = make([]int, m)
		for j := 0; j < m; j++ {
			memo[i][j] = -1
		}
	}

	return memo
}

func main() {
	type Input struct {
		words  []string
		target string
	}

	testCases := []struct {
		input    Input
		expected int
	}{
		{
			Input{
				[]string{"acca", "bbbb", "caca"},
				"aba",
			},
			6,
		},
		{
			Input{
				[]string{"abba", "baab"},
				"bab",
			},
			4,
		},
		{
			Input{
				[]string{
					"cbabddddbc", "addbaacbbd", "cccbacdccd", "cdcaccacac", "dddbacabbd",
					"bdbdadbccb", "ddadbacddd", "bbccdddadd", "dcabaccbbd", "ddddcddadc",
					"bdcaaaabdd", "adacdcdcdd", "cbaaadbdbb", "bccbabcbab", "accbdccadd",
					"dcccaaddbc", "cccccacabd", "acacdbcbbc", "dbbdbaccca", "bdbddbddda",
					"daabadbacb", "baccdbaada", "ccbabaabcb", "dcaabccbbb", "bcadddaacc",
					"acddbbdccb", "adbddbadab", "dbbcdcbcdd", "ddbabbadbb", "bccbcbbbab",
					"dabbbdbbcb", "dacdabadbb", "addcbbabab", "bcbbccadda", "abbcacadac",
					"ccdadcaada", "bcacdbccdb",
				},
				"bcbbcccc",
			},
			677452090,
		},
	}

	for _, tc := range testCases {
		actual := numWays(tc.input.words, tc.input.target)
		if tc.expected != actual {
			fmt.Printf("numWays(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
