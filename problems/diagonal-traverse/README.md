# 498. Diagonal Traverse

- Difficulty: Medium
- Topics: Array, Matrix, Simulation
- Link: https://leetcode.com/problems/diagonal-traverse/

## Description

Given an `m x n` matrix `mat`, return *an array of all the elements of the array in a diagonal order*.

**Example 1:**

![](https://assets.leetcode.com/uploads/2021/04/10/diag1-grid.jpg)

```
Input: mat = [[1,2,3],[4,5,6],[7,8,9]]
Output: [1,2,4,7,5,3,6,8,9]
```

**Example 2:**

```
Input: mat = [[1,2],[3,4]]
Output: [1,2,3,4]
```

**Constraints:**

- `m == mat.length`
- `n == mat[i].length`
- `1 <= m, n <= 10^4`
- `1 <= m * n <= 10^4`
- `-10^5 <= mat[i][j] <= 10^5`

## Solution

During diagonal traversal, we alternate between moving up-right and down-left, reversing direction when we reach a matrix edge.

Take care of edge cases:

- When we hit the top-right, move down instead of right.
- When we hit the bottom-left, move right instead of down.
