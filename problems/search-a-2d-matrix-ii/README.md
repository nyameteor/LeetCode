# 240. Search a 2D Matrix II

- Difficulty: Medium
- Topics: Array, Binary Search, Divide and Conquer, Matrix
- Link: https://leetcode.com/problems/search-a-2d-matrix-ii/

## Description

Write an efficient algorithm that searches for a value `target` in an `m x n` integer matrix `matrix`. This matrix has the following properties:

- Integers in each row are sorted in ascending from left to right.
- Integers in each column are sorted in ascending from top to bottom.

**Example 1:**

![](https://assets.leetcode.com/uploads/2020/11/24/searchgrid2.jpg)

```
Input: matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 5
Output: true
```

**Example 2:**

![](https://assets.leetcode.com/uploads/2020/11/24/searchgrid.jpg)

```
Input: matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 20
Output: false
```

**Constraints:**

- `m == matrix.length`
- `n == matrix[i].length`
- `1 <= n, m <= 300`
- `-10^9 <= matrix[i][j] <= 10^9`
- All the integers in each row are **sorted** in ascending order.
- All the integers in each column are **sorted** in ascending order.
- `-10^9 <= target <= 10^9`

## Solution

### Approach: Greedy

Start from the **top-right corner** and move:

- **Left** if the value is **greater** than `target`.
- **Down** if the value is **smaller** than `target`.
- **Return `true`** if a match is found.

**Time Complexity**:

- Worst case: `O(m + n)`, eliminating a row or column at each step.

### Approach: Binary Search on Rows

Use binary search to locate the column boundary more efficiently instead of scanning row-by-row.

- Start from the **top-right** as before.
- If `matrix[row][col] < target`, perform binary search in that row to find the first element >= `target`.

**Time Complexity**:

- Worst case: `O(m log n)`, which is better when `n >> m`.
