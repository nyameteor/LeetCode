package main

import (
	"fmt"
	"math"
)

func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)

	totalGas := 0
	minGas := math.MaxInt32
	minIdx := -1

	for i := 0; i < n; i++ {
		totalGas += gas[i]
		totalGas -= cost[i]
		if totalGas < minGas {
			minGas = totalGas
			minIdx = i
		}
	}

	if totalGas < 0 {
		return -1
	}

	return (minIdx + 1) % n
}

func main() {
	type Input struct {
		gas  []int
		cost []int
	}

	testCases := []struct {
		input    Input
		expected int
	}{
		{
			input: Input{
				gas:  []int{1, 2, 3, 4, 5},
				cost: []int{3, 4, 5, 1, 2},
			},
			expected: 3,
		},
		{
			input: Input{
				gas:  []int{2, 3, 4},
				cost: []int{3, 4, 3},
			},
			expected: -1,
		},
		{
			input: Input{
				gas:  []int{5, 8, 2, 8},
				cost: []int{6, 5, 6, 6},
			},
			expected: 3,
		},
		{
			input: Input{
				gas:  []int{3, 1, 1},
				cost: []int{1, 2, 2},
			},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		actual := canCompleteCircuit(tc.input.gas, tc.input.cost)
		if actual != tc.expected {
			fmt.Printf("canCompleteCircuit(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
