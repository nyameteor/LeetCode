package main

import "fmt"

// Definition for a QuadTree node.
type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func construct(grid [][]int) *Node {
	return build(grid, 0, 0, len(grid))
}

func build(grid [][]int, x, y, size int) *Node {
	if check(grid, x, y, size) {
		return &Node{
			Val:    grid[x][y] == 1,
			IsLeaf: true,
		}
	}

	half := size / 2
	return &Node{
		Val:         false,
		IsLeaf:      false,
		TopLeft:     build(grid, x, y, half),
		TopRight:    build(grid, x, y+half, half),
		BottomLeft:  build(grid, x+half, y, half),
		BottomRight: build(grid, x+half, y+half, half),
	}
}

func check(grid [][]int, x, y, size int) bool {
	val := grid[x][y]
	for i := x; i < x+size; i++ {
		for j := y; j < y+size; j++ {
			if grid[i][j] != val {
				return false
			}
		}
	}
	return true
}

func main() {
	testCases := []struct {
		input    [][]int
		expected *Node
	}{
		{
			input: [][]int{{0, 1}, {1, 0}},
			expected: &Node{
				Val:         false,
				IsLeaf:      false,
				TopLeft:     &Node{Val: false, IsLeaf: true},
				TopRight:    &Node{Val: true, IsLeaf: true},
				BottomLeft:  &Node{Val: true, IsLeaf: true},
				BottomRight: &Node{Val: false, IsLeaf: true},
			},
		},
		{
			input: [][]int{{0}},
			expected: &Node{
				Val:    false,
				IsLeaf: true,
			},
		},
	}

	var treesEqual func(a, b *Node) bool
	treesEqual = func(a, b *Node) bool {
		if a == nil && b == nil {
			return true
		}
		if a == nil || b == nil {
			return false
		}
		return a.Val == b.Val && a.IsLeaf == b.IsLeaf &&
			treesEqual(a.TopLeft, b.TopLeft) &&
			treesEqual(a.TopRight, b.TopRight) &&
			treesEqual(a.BottomLeft, b.BottomLeft) &&
			treesEqual(a.BottomRight, b.BottomRight)
	}

	for _, tc := range testCases {
		actual := construct(tc.input)
		if !treesEqual(actual, tc.expected) {
			fmt.Printf("construct(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
