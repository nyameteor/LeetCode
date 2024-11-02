package main

import (
	"fmt"
	"sort"
)

func minimumTotalDistance(robot []int, factory [][]int) int64 {
	sort.Ints(robot)
	sort.Slice(factory, func(i, j int) bool {
		return factory[i][0] < factory[j][0]
	})

	factoryPos := []int{}
	for _, f := range factory {
		for range f[1] {
			factoryPos = append(factoryPos, f[0])
		}
	}

	m, n := len(robot), len(factoryPos)

	memo := make([][]int64, m)
	for i := range memo {
		memo[i] = make([]int64, n)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	return calcMinDistance(robot, factoryPos, memo, 0, 0)
}

func calcMinDistance(robot []int, factoryPos []int, memo [][]int64, robotIdx, factoryIdx int) int64 {
	if robotIdx >= len(robot) {
		return 0
	}
	if factoryIdx >= len(factoryPos) {
		return 1 << 62
	}

	if memo[robotIdx][factoryIdx] != -1 {
		return memo[robotIdx][factoryIdx]
	}

	var dist1, dist2 int64
	dist1 = int64(abs(robot[robotIdx]-factoryPos[factoryIdx])) + calcMinDistance(robot, factoryPos, memo, robotIdx+1, factoryIdx+1)
	dist2 = calcMinDistance(robot, factoryPos, memo, robotIdx, factoryIdx+1)

	memo[robotIdx][factoryIdx] = min(dist1, dist2)
	return memo[robotIdx][factoryIdx]
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func main() {
	type Input struct {
		robot   []int
		factory [][]int
	}
	testCases := []struct {
		input    Input
		expected int64
	}{
		{
			input: Input{
				robot: []int{0, 4, 6},
				factory: [][]int{
					{2, 2},
					{6, 2},
				},
			},
			expected: 4,
		},
		{
			input: Input{
				robot: []int{1, -1},
				factory: [][]int{
					{-2, 1},
					{2, 1},
				},
			},
			expected: 2,
		},
	}

	for _, tc := range testCases {
		res := minimumTotalDistance(tc.input.robot, tc.input.factory)
		if res != int64(tc.expected) {
			fmt.Printf("minimumTotalDistance(%v, %v) = %v, expected = %v\n", tc.input.robot, tc.input.factory, res, tc.expected)
		}
	}
}
