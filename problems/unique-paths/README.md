# 62. Unique Paths

- Difficulty: Medium
- Topics: Math, Dynamic Programming, Combinatorics
- Link: https://leetcode.com/problems/unique-paths/

## Description

A robot is located at the top-left corner of a `m x n` grid (marked 'Start' in the diagram below).

The robot can only move either down or right at any point in time. The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).

How many possible unique paths are there?

**Example 1:**

![img](https://assets.leetcode.com/uploads/2018/10/22/robot_maze.png)

```
Input: m = 3, n = 7
Output: 28
```

**Example 2:**

```
Input: m = 3, n = 2
Output: 3
Explanation:
From the top-left corner, there are a total of 3 ways to reach the bottom-right corner:
1. Right -> Down -> Down
2. Down -> Down -> Right
3. Down -> Right -> Down
```

**Example 3:**

```
Input: m = 7, n = 3
Output: 28
```

**Example 4:**

```
Input: m = 3, n = 3
Output: 6
```

**Constraints:**

- `1 <= m, n <= 100`
- It's guaranteed that the answer will be less than or equal to `2 * 10^9`.

## Solution

### Observations

- The robot can only move **down** or **right**.
- Any path is a combination of `m-1` downs and `n-1` rights.
- In the DP approach, each cell depends on the cell **below** and the cell **to the right** (top-down), or the cell **above** and the cell **to the left** (bottom-up).
- The base case is the destination cell, which has **1** path (staying there).
