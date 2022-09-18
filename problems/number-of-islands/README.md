# 200. Number of Islands

- Difficulty: Medium
- Topics: Array, Depth-First Search, Breadth-First Search, Union Find, Matrix
- Link: https://leetcode.com/problems/number-of-islands/

## Description

Given an `m x n` 2D binary grid `grid` which represents a map of `'1'`s (land) and `'0'`s (water), return _the number of islands_.

An **island** is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.

**Example 1:**

```
Input: grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
Output: 1
```

**Example 2:**

```
Input: grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
Output: 3
```

**Constraints:**

- `m == grid.length`
- `n == grid[i].length`
- `1 <= m, n <= 300`
- `grid[i][j]` is `'0'` or `'1'`.

## Solution

### Union Find

矩阵中的某个元素 a 和 b 是否连接代表一种等价关系（Equivalence Relations）或连通关系，适合使用并查集来处理。

- 初始化集合 s，s 的大小为 `m * n`，用 `i * col + j` 标识 matrix 中的每个元素
- 遍历矩阵获得连通关系，合并连通的元素为一个集合
- 遍历集合，不交集（不交集是一个单独的连通集合）数量 = root 为自身的元素数量
- 岛屿数量 = 不交集数量 - 水的数量

Todo: 规范化表述

### Depth-First Search

Todo
