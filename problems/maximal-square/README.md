# 221. Maximal Square

- Difficulty: Medium
- Topics: Array, Dynamic Programming, Matrix
- Link: https://leetcode.com/problems/maximal-square/

## Description

Given an `m x n` binary `matrix` filled with `0`'s and `1`'s, _find the largest square containing only_ `1`'s _and return its area_.

**Example 1:**

![img](https://assets.leetcode.com/uploads/2020/11/26/max1grid.jpg)

```
Input: matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
Output: 4
```

**Example 2:**

![img](https://assets.leetcode.com/uploads/2020/11/26/max2grid.jpg)

```
Input: matrix = [["0","1"],["1","0"]]
Output: 1
```

**Example 3:**

```
Input: matrix = [["0"]]
Output: 0
```

**Constraints:**

- `m == matrix.length`
- `n == matrix[i].length`
- `1 <= m, n <= 300`
- `matrix[i][j]` is `'0'` or `'1'`.

## Solution

### Approach: Dynamic Programming

The key observation is:

1. If `matrix[i][j]` is `'0'`, it cannot contribute to any square.
2. If `matrix[i][j]` is `'1'`, the largest square ending at `(i, j)` depends on its top, left, and top-left diagonal neighbors:

   ```
   dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
   ```

3. The final answer is the **square of the largest side length** found in `dp`.

This approach ensures `O(m * n)` time complexity.

### References

https://leetcode.com/problems/maximal-square/solutions/600149/python-thinking-process-diagrams-dp-approach/
