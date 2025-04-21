# 441. Arranging Coins

- Difficulty: Easy
- Topics: Math, Binary Search
- Link: https://leetcode.com/problems/arranging-coins/

## Description

You have `n` coins and you want to build a staircase with these coins. The staircase consists of `k` rows where the `ith` row has exactly `i` coins. The last row of the staircase **may be** incomplete.

Given the integer `n`, return _the number of **complete rows** of the staircase you will build_.

**Example 1:**

![img](https://assets.leetcode.com/uploads/2021/04/09/arrangecoins1-grid.jpg)

```
Input: n = 5
Output: 2
Explanation: Because the 3rd row is incomplete, we return 2.
```

**Example 2:**

![img](https://assets.leetcode.com/uploads/2021/04/09/arrangecoins2-grid.jpg)

```
Input: n = 8
Output: 3
Explanation: Because the 4th row is incomplete, we return 3.
```

**Constraints:**

- `1 <= n <= 2^31 - 1`

## Solution

### Approach: Binary Search

- Each `i-th` row requires `i` coins, forming a total of `k * (k + 1) / 2` coins for `k` full rows.
- We want the largest `k` such that: `k * (k + 1) / 2 <= n`.
- Use **binary search** to efficiently find the maximum valid `k`.

### References

- [[C++/Java/Python] O(sqrt(n)), O(logn), O(1) Approaches with Explanation](https://leetcode.com/problems/arranging-coins/solutions/1559984/C++JavaPython-O/)
