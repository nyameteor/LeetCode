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

本题是一道模拟题。螺旋遍历矩阵是个向内收缩的过程，遍历时的上下界会随之变化。

```shell
matrix
c               n
1   2   3   4   r
5   6   7   8
9   10  11  12
                m

top
c               n
1   2   3   4   r       [c, n), r++
*   *   *   *
*   *   *   *
                m

right
c               n
*   *   *   *
*   *   *   8   r
*   *   *   12          [r, m), n--
                m

bottom
c           n
*   *   *   *
*   *   *   *   r
9   10  11  *           [n-1, c), m--
                m

left
c           n
*   *   *   *
5   *   *   *   r
*   *   *   *   m       [m-1, r), c++

top
    c       n
*   *   *   *
*   6   7   *   r       [c, n), r++
*   *   *   *   m
...
```

另外当 matrix 只有一行或一列时（如 {1, 2} 和 {{1}, {2}}），需要额外添加条件防止 bottom 和 left 遍历到重复的内容：

- 若 matrix 只有一行（ r == m ），则不遍历 bottom
- 若 matrix 只有一列（ c == n ），则不遍历 left
