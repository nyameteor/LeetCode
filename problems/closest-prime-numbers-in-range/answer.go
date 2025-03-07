package main

import (
	"fmt"
	"slices"
)

func closestPrimes(left int, right int) []int {
	primes := sieveOfEratosthenes(right)
	res := []int{-1, -1}

	minDistance := right
	prevPrime := -1
	for i := left; i <= right; i++ {
		if !primes[i] {
			continue
		}
		if prevPrime != -1 && i-prevPrime < minDistance {
			minDistance = i - prevPrime
			res[0], res[1] = prevPrime, i
		}
		prevPrime = i
	}

	return res
}

func sieveOfEratosthenes(n int) []bool {
	primes := make([]bool, n+1)

	for i := 2; i <= n; i++ {
		primes[i] = true
	}

	for i := 2; i*i <= n; i++ {
		if primes[i] {
			for j := i * i; j <= n; j += i {
				primes[j] = false
			}
		}
	}

	return primes
}

func main() {
	type Input struct {
		left  int
		right int
	}

	testCases := []struct {
		input    Input
		expected []int
	}{
		{
			input: Input{
				left:  10,
				right: 19,
			},
			expected: []int{11, 13},
		},
		{
			input: Input{
				left:  4,
				right: 6,
			},
			expected: []int{-1, -1},
		},
		{
			input: Input{
				left:  19,
				right: 31,
			},
			expected: []int{29, 31},
		},
		{
			input: Input{
				left:  710119,
				right: 710189,
			},
			expected: []int{710119, 710189},
		},
		{
			input: Input{
				left:  1,
				right: 1000000,
			},
			expected: []int{2, 3},
		},
	}

	for _, tc := range testCases {
		actual := closestPrimes(tc.input.left, tc.input.right)
		if !slices.Equal(actual, tc.expected) {
			fmt.Printf("closestPrimes(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
