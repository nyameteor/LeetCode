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

### Binary Search

按照条件，可以使用两次二分查找，先搜列，再搜行。列的匹配需要加额外判断条件。

```shell
matrix              target = 3
   0  1  2  3
0  1  3  5  7
1  10 11 16 20
2  23 30 34 60

search in row
0   1   2
1   10  23          target < matrix[m][0], r = m - 1
l  m  h

1   10  23          m < rowSize - 1, matrix[m][0] <= target < matrix[m+1][0], row = m
l,m,r

search in col
0  1  2  3
1  3  5  7
(standar binary search, skip...)

row = 0, col = 1
```
