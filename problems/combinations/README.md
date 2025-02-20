# 77. Combinations

- Difficulty: Medium
- Topics: Backtracking
- Link: https://leetcode.com/problems/combinations/

## Description

Given two integers `n` and `k`, return *all possible combinations of* `k` *numbers chosen from the range* `[1, n]`.

You may return the answer in **any order**.

**Example 1:**

```
Input: n = 4, k = 2
Output: [[1,2],[1,3],[1,4],[2,3],[2,4],[3,4]]
Explanation: There are 4 choose 2 = 6 total combinations.
Note that combinations are unordered, i.e., [1,2] and [2,1] are considered to be the same combination.

```

**Example 2:**

```
Input: n = 1, k = 1
Output: [[1]]
Explanation: There is 1 choose 1 = 1 total combination.

```

**Constraints:**

- `1 <= n <= 20`
- `1 <= k <= n`

## Solution

1. **Backtracking Approach**: We use DFS to explore all possible combinations of `k` numbers.  
2. **Base Case**: When `path` reaches length `k`, we add a copy to `res`.  
3. **Recursive Exploration**: We iterate from `start` to `n`, adding each number to `path` and moving to the next.  
4. **Pruning**: By starting from `i+1`, we ensure elements are chosen in increasing order, avoiding duplicates.
