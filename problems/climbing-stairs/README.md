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

### Key Idea

This problem is a classic example of **dynamic programming**. The number of ways to reach the `n`-th step depends on the previous two steps:

- `dp[i] = dp[i-1] + dp[i-2]`

### Observation

- This is similar to the **Fibonacci sequence**.
- Space can be optimized by storing only the last two values, reducing space complexity to **O(1)**.
