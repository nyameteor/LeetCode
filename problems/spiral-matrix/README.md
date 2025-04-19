# 54. Spiral Matrix

- Difficulty: Medium
- Topics: Array, Matrix, Simulation
- Link: https://leetcode.com/problems/spiral-matrix/

## Description

Given an `m x n` `matrix`, return _all elements of the_ `matrix` _in spiral order_.

**Example 1:**

![img](https://assets.leetcode.com/uploads/2020/11/13/spiral1.jpg)

```
Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [1,2,3,6,9,8,7,4,5]
```

**Example 2:**

![img](https://assets.leetcode.com/uploads/2020/11/13/spiral.jpg)

```
Input: matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
Output: [1,2,3,4,8,12,11,10,9,5,6,7]
```

**Constraints:**

- `m == matrix.length`
- `n == matrix[i].length`
- `1 <= m, n <= 10`
- `-100 <= matrix[i][j] <= 100`

## Solution

### Key Idea

Use four boundaries (`top`, `bottom`, `left`, `right`) to simulate the spiral traversal. At each step, move in one of the four directions:

- Left to right
- Top to bottom
- Right to left (only if `top <= bottom`)
- Bottom to top (only if `left <= right`)

After traversing each direction, shrink the boundary to avoid revisiting. Repeat until all elements are visited.
