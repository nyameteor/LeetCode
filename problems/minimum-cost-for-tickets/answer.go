package main

import "fmt"

func mincostTickets(days []int, costs []int) int {
	memo := map[int]int{}
	return helper(days, costs, 0, memo)
}

func helper(days []int, costs []int, i int, memo map[int]int) int {
	if v, ok := memo[i]; ok {
		return v
	}

	if i >= len(days) {
		return 0
	}

	res1 := costs[0] + helper(days, costs, findNextDayIndex(days, i, 1), memo)
	res2 := costs[1] + helper(days, costs, findNextDayIndex(days, i, 7), memo)
	res3 := costs[2] + helper(days, costs, findNextDayIndex(days, i, 30), memo)

	res := min(res1, res2, res3)

	memo[i] = res

	return res
}

func findNextDayIndex(days []int, i int, pass int) int {
	j := i
	for j < len(days) && days[j] < days[i]+pass {
		j++
	}
	return j
}

func main() {
	type Input struct {
		days  []int
		costs []int
	}

	testCases := []struct {
		input    Input
		expected int
	}{
		{
			Input{
				[]int{1, 4, 6, 7, 8, 20},
				[]int{2, 7, 15},
			},
			11,
		},
		{
			Input{
				[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31},
				[]int{2, 7, 15},
			},
			17,
		},
	}

	for _, tc := range testCases {
		actual := mincostTickets(tc.input.days, tc.input.costs)
		if actual != tc.expected {
			fmt.Printf("mincostTickets(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
