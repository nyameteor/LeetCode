# 518. Coin Change II

- Difficulty: Medium
- Topics: Array, Dynamic Programming
- Link: https://leetcode.com/problems/coin-change-ii/

## Description

You are given an integer array `coins` representing coins of different denominations and an integer `amount` representing a total amount of money.

Return *the number of combinations that make up that amount*. If that amount of money cannot be made up by any combination of the coins, return `0`.

You may assume that you have an infinite number of each kind of coin.

The answer is **guaranteed** to fit into a signed **32-bit** integer.

**Example 1:**

```
Input: amount = 5, coins = [1,2,5]
Output: 4
Explanation: there are four ways to make up the amount:
5=5
5=2+2+1
5=2+1+1+1
5=1+1+1+1+1
```

**Example 2:**

```
Input: amount = 3, coins = [2]
Output: 0
Explanation: the amount of 3 cannot be made up just with coins of 2.
```

**Example 3:**

```
Input: amount = 10, coins = [10]
Output: 1
```

**Constraints:**

- `1 <= coins.length <= 300`
- `1 <= coins[i] <= 5000`
- All the values of `coins` are **unique**.
- `0 <= amount <= 5000`

## Solution

## Observations

- **Problem Breakdown**: The goal is to count the number of ways to form an amount using the given coins. The order of coins doesn't matter (combinations, not permutations).
- **Recursive Structure**: We can recursively break down the problem by including or excluding a coin at each step.

## Intuition

### Top-Down (DFS with Memoization)

- **Subproblems**: `dfs(remain, i)` counts ways to make up `remain` using coins starting from index `i`.
- **Choice**: For each coin, we can either include it (subtract from remaining amount) or exclude it. We allow multiple uses of the same coin, so we continue from the current or next coin.
- **Memoization**: Saves results of subproblems to avoid redundant calculations.

### Bottom-Up (Tabulation)

- **State Representation**: `dp[i]` is the number of ways to make amount `i`.
- **Initialization**: `dp[0] = 1` (one way to make `0` - no coins).
- **State Transition**: For each coin, update `dp[i]` by adding `dp[i-coin]`.
- **Final Answer**: `dp[amount]` holds the result.

## Complexity

- **Time**: Both approaches have `O(amount * n)` time complexity.
- **Space**: Top-Down uses `O(amount * n)` space (memoization), while Bottom-Up uses `O(amount)` space (`dp` array).
