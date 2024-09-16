package main

import (
	"fmt"
	"sort"
	"strconv"
)

func findMinDifference(timePoints []string) int {
	times := make([]int, len(timePoints))
	for i, timePoint := range timePoints {
		hour, err := strconv.Atoi(timePoint[0:2])
		if err != nil {
			return -1
		}
		minute, err := strconv.Atoi(timePoint[3:5])
		if err != nil {
			return -1
		}
		times[i] = hour*60 + minute
	}

	sort.Ints(times)

	minDiff := 24 * 60
	for i := 1; i < len(times); i++ {
		diff := times[i] - times[i-1]
		if minDiff > diff {
			minDiff = diff
		}
		if minDiff == 0 {
			return minDiff
		}
	}
	diff := times[0] - times[len(times)-1] + 24*60
	if minDiff > diff {
		minDiff = diff
	}

	return minDiff
}

func main() {
	var timePoints []string
	var res int

	timePoints = []string{"23:59", "00:00"}
	res = findMinDifference(timePoints)
	// 1
	fmt.Println(res)

	timePoints = []string{"00:00", "23:59", "00:00"}
	res = findMinDifference(timePoints)
	// 0
	fmt.Println(res)
}
