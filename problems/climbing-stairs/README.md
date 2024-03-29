# 70. Climbing Stairs

- Difficulty: Easy
- Topics: Math, Dynamic Programming, Memoization
- Link: https://leetcode.com/problems/climbing-stairs/

## Description

You are climbing a staircase. It takes `n` steps to reach the top.

Each time you can either climb `1` or `2` steps. In how many distinct ways can you climb to the top?

**Example 1:**

```
Input: n = 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps
```

**Example 2:**

```
Input: n = 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step
```

**Constraints:**

- `1 <= n <= 45`

## Solution

每次可以爬 1 或 2 步任选，是无界背包问题，要求遍历整颗树，返回所有符合要求的情况的数量。

用 DP 求解：

- 基本情况：`n == 0, f(n) == 1; n < 0, f(n) == 0;`
- 递推关系：`f(n) = f(n-1) + f(n-2);`
- 添加 memo 优化为线性时间
- time: `O(n)`
- space: `O(n)`
