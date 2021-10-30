# 994. Rotting Oranges

- Difficulty: Medium
- Topics: Array, Breadth-First Search, Matrix
- Link: https://leetcode.com/problems/rotting-oranges/

## Description

You are given an `m x n` `grid` where each cell can have one of three values:

- `0` representing an empty cell,
- `1` representing a fresh orange, or
- `2` representing a rotten orange.

Every minute, any fresh orange that is **4-directionally adjacent** to a rotten orange becomes rotten.

Return _the minimum number of minutes that must elapse until no cell has a fresh orange_. If _this is impossible, return_ `-1`.

**Example 1:**

![img](https://assets.leetcode.com/uploads/2019/02/16/oranges.png)

```
Input: grid = [[2,1,1],[1,1,0],[0,1,1]]
Output: 4
```

**Example 2:**

```
Input: grid = [[2,1,1],[0,1,1],[1,0,1]]
Output: -1
Explanation: The orange in the bottom left corner (row 2, column 0) is never rotten, because rotting only happens 4-directionally.
```

**Example 3:**

```
Input: grid = [[0,2]]
Output: 0
Explanation: Since there are already no fresh oranges at minute 0, the answer is just 0.
```

**Constraints:**

- `m == grid.length`
- `n == grid[i].length`
- `1 <= m, n <= 10`
- `grid[i][j]` is `0`, `1`, or `2`.

## Solution

### Breadth-First Search

1. 将所有初始时烂橘子的位置推入队列； 开始时计算所有新鲜橙子；
2. 从烂橙子遍历，弹出队列头，到相邻的 4 个位置（上下左右），如果橙子是新鲜的，则让它腐烂，并减少新鲜橙子的数量；
3. 重复步骤 2 后增加`time`，直到队列为空；
4. 如果还有新鲜的橙子就返回-1，否则返回`time`。

BFS 时每个节点有四个分支（上下左右），可以使用一个 pair 数组记录其移动方向，使四个分支可以在循环中处理。
