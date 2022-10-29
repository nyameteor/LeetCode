package main

import "fmt"

type point struct {
	x int
	y int
}

type vector struct {
	x int
	y int
}

func largestOverlap(img1 [][]int, img2 [][]int) int {
	maxOverlaps := 0

	points1 := nonZeroPoints(img1)
	points2 := nonZeroPoints(img2)

	groupCount := map[vector]int{}

	for _, p1 := range points1 {
		for _, p2 := range points2 {
			v := vector{p1.x - p2.x, p1.y - p2.y}
			groupCount[v]++
			maxOverlaps = max(maxOverlaps, groupCount[v])
		}
	}

	return maxOverlaps
}

func nonZeroPoints(img [][]int) []point {
	points := []point{}
	n := len(img)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if img[i][j] == 1 {
				points = append(points, point{i, j})
			}
		}
	}
	return points
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func main() {
	var img1 [][]int
	var img2 [][]int

	img1 = [][]int{{1, 1, 0}, {0, 1, 0}, {0, 1, 0}}
	img2 = [][]int{{0, 0, 0}, {0, 1, 1}, {0, 0, 1}}
	fmt.Println(largestOverlap(img1, img2) == 3)

	img1 = [][]int{{1}}
	img2 = [][]int{{1}}
	fmt.Println(largestOverlap(img1, img2) == 1)
}
