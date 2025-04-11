package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Definition for a Node.
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	return bfsOptimized(root)
}

func bfsOptimized(root *Node) *Node {
	level := root
	for level != nil {
		cur := level
		for cur != nil {
			if cur.Left != nil {
				cur.Left.Next = cur.Right
				if cur.Next != nil {
					cur.Right.Next = cur.Next.Left
				}
			}
			cur = cur.Next
		}
		level = level.Left
	}
	return root
}

func bfs(root *Node) *Node {
	if root == nil {
		return nil
	}

	queue := []*Node{root}
	for len(queue) > 0 {
		n := len(queue)
		var right *Node
		for range n {
			cur := queue[0]
			queue = queue[1:]

			cur.Next = right
			right = cur

			if cur.Right != nil {
				queue = append(queue, cur.Right)
				queue = append(queue, cur.Left)
			}
		}
	}

	return root
}

func main() {
	testCases := []struct {
		input *Node
	}{
		{
			input: &Node{
				Val: 1,
				Left: &Node{
					Val:   2,
					Left:  &Node{Val: 4},
					Right: &Node{Val: 5},
				},
				Right: &Node{
					Val:   3,
					Left:  &Node{Val: 6},
					Right: &Node{Val: 7},
				},
			},
		},
		{
			input: nil,
		},
	}

	solutions := []struct {
		name string
		fn   func(*Node) *Node
	}{
		{
			name: "Default Solution",
			fn:   connect,
		},
		{
			name: "BFS",
			fn:   bfs,
		},
		{
			name: "BFS Optimized",
			fn:   bfsOptimized,
		},
	}

	treeWithNextString := func(root *Node) string {
		var sb strings.Builder
		level := root

		for level != nil {
			cur := level
			for cur != nil {
				sb.WriteString(strconv.Itoa(cur.Val))
				if cur.Next != nil {
					sb.WriteString(" -> ")
				}
				cur = cur.Next
			}
			sb.WriteString(" -> nil\n")
			level = level.Left
		}

		return sb.String()
	}

	for _, solution := range solutions {
		for _, tc := range testCases {
			actual := solution.fn(tc.input)
			fmt.Printf("Solution Name: %v\n", solution.name)
			fmt.Printf("Input Tree: %v\n", tc.input)
			fmt.Println("Next Right Pointers:")
			fmt.Println(treeWithNextString(actual))
		}
	}
}
