package main

import "fmt"

func canReach(arr []int, start int) bool {
	n := len(arr)

	if n == 0 {
		return false
	}

	seen := make([]bool, n)
	seen[start] = true
	queue := []int{start}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if arr[cur] == 0 {
			return true
		}

		left := cur - arr[cur]
		if left >= 0 && left < n && !seen[left] {
			seen[left] = true
			queue = append(queue, left)
		}

		right := cur + arr[cur]
		if right >= 0 && right < n && !seen[right] {
			seen[right] = true
			queue = append(queue, right)
		}
	}

	return false
}

func main() {
	type Input struct {
		arr   []int
		start int
	}

	testCases := []struct {
		input    *Input
		expected bool
	}{
		{
			input: &Input{
				arr:   []int{4, 2, 3, 0, 3, 1, 2},
				start: 5,
			},
			expected: true,
		},
		{
			input: &Input{
				arr:   []int{4, 2, 3, 0, 3, 1, 2},
				start: 0,
			},
			expected: true,
		},
		{
			input: &Input{
				arr:   []int{3, 0, 2, 1, 2},
				start: 2,
			},
			expected: false,
		},
		{
			input: &Input{
				arr:   []int{0, 3, 0, 6, 3, 3, 4},
				start: 6,
			},
			expected: true,
		},
	}

	for _, tc := range testCases {
		actual := canReach(tc.input.arr, tc.input.start)
		if actual != tc.expected {
			fmt.Printf("canReach(%v) = %v, expected = %v\n", tc.input, actual, tc.expected)
		}
	}
}
