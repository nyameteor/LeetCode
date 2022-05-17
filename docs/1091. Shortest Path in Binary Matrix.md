# 1091. Shortest Path in Binary Matrix

- Difficulty: Medium
- Topics: Array, Breadth-First Search, Matrix
- Link: https://leetcode.com/problems/shortest-path-in-binary-matrix/

## Description

Given an `n x n` binary matrix `grid`, return _the length of the shortest **clear path** in the matrix_. If there is no clear path, return `-1`.

A **clear path** in a binary matrix is a path from the **top-left** cell (i.e., `(0, 0)`) to the **bottom-right** cell (i.e., `(n - 1, n - 1)`) such that:

- All the visited cells of the path are `0`.
- All the adjacent cells of the path are **8-directionally** connected (i.e., they are different and they share an edge or a corner).

The **length of a clear path** is the number of visited cells of this path.

**Example 1:**

![img](https://assets.leetcode.com/uploads/2021/02/18/example1_1.png)

```
Input: grid = [[0,1],[1,0]]
Output: 2
```

**Example 2:**

![img](https://assets.leetcode.com/uploads/2021/02/18/example2_1.png)

```
Input: grid = [[0,0,0],[1,1,0],[1,1,0]]
Output: 4
```

**Example 3:**

```
Input: grid = [[1,0,0],[1,1,0],[1,1,0]]
Output: -1
```

**Constraints:**

- `n == grid.length`
- `n == grid[i].length`
- `1 <= n <= 100`
- `grid[i][j] is 0 or 1`

## Solution

### Breadth-First Search

Use BFS approach to find the shortest path.

Denote `N` as the width and height of `grid`.

- Use `visited` to record if cell is visited.
- To track levels, use `len` to record number of nodes in current level.
- Start at cell `(0, 0)`, search all adjacent cells:
  - if the position of cell `(i, j)` is valid, and its value equal to `0`, and has not visited, then push `(i, j)` to the queue.
- When we reach at cell `(N - 1, N - 1)`, return current level.
- Cannot find, return `-1`.
