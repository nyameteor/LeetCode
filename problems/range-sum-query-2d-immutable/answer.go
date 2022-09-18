package main

type NumMatrix struct {
	preSum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	if len(matrix) == 0 {
		return NumMatrix{nil}
	}

	m := len(matrix)
	n := len(matrix[0])

	preSum := make([][]int, m+1)
	preSum[0] = make([]int, n+1)

	for i := 0; i < m; i++ {
		preSum[i+1] = make([]int, n+1)
		for j := 0; j < n; j++ {
			preSum[i+1][j+1] = preSum[i][j+1] + preSum[i+1][j] - preSum[i][j] + matrix[i][j]
		}
	}

	return NumMatrix{preSum}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	preSum := this.preSum
	return preSum[row2+1][col2+1] - preSum[row1][col2+1] - preSum[row2+1][col1] + preSum[row1][col1]
}
