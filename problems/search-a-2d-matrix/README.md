# 74. Search a 2D Matrix

- Difficulty: Medium
- Topics: Array, Binary Search, Matrix
- Link: https://leetcode.com/problems/search-a-2d-matrix/

## Description

Write an efficient algorithm that searches for a value in an `m x n` matrix. This matrix has the following properties:

- Integers in each row are sorted from left to right.
- The first integer of each row is greater than the last integer of the previous row.

**Example 1:**

![img](https://assets.leetcode.com/uploads/2020/10/05/mat.jpg)

```
Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
Output: true
```

**Example 2:**

![img](https://assets.leetcode.com/uploads/2020/10/05/mat2.jpg)

```
Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
Output: false
```

**Constraints:**

- `m == matrix.length`
- `n == matrix[i].length`
- `1 <= m, n <= 100`
- `-10^4 <= matrix[i][j], target <= 10^4`

## Solution

### Approach: Binary Search on Flattened Matrix

Treat the matrix as a 1D array and apply binary search:

1. Convert the 2D indices to a single index: `row = m / numCols`, `col = m % numCols`.
2. Binary search:
   - If the value equals the target, return `true`.
   - If greater, adjust `r = m - 1`; if smaller, adjust `l = m + 1`.
3. Return `false` if the target is not found.

**Complexity**:

- **Time**: `O(log(m * n))`
- **Space**: `O(1)`
