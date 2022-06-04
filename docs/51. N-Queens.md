# 51. N-Queens

- Difficulty: Hard
- Topics: Array, Backtracking
- Link: https://leetcode.com/problems/n-queens/

## Description

The **n-queens** puzzle is the problem of placing `n` queens on an `n x n` chessboard such that no two queens attack each other.

Given an integer `n`, return \*all distinct solutions to the **n-queens puzzle\***. You may return the answer in **any order**.

Each solution contains a distinct board configuration of the n-queens' placement, where `'Q'` and `'.'` both indicate a queen and an empty space, respectively.

**Example 1:**

![img](https://assets.leetcode.com/uploads/2020/11/13/queens.jpg)

```
Input: n = 4
Output: [[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
Explanation: There exist two distinct solutions to the 4-queens puzzle as shown above
```

**Example 2:**

```
Input: n = 1
Output: [["Q"]]
```

**Constraints:**

- `1 <= n <= 9`

## Solution

### Backtracking

We can use depth-first search approach to search all permutations of position for `N` queens.

During the search:

- if `size(B) == N`, then we find a valid permutation, add to result, and exit current search.
- if a position `{i, j}` is valid(no two queen attack each other), then add to board `B`.
- search positions in next row `{i+1, 0}, {i+1, 1}, ..., {i+1, N-1}`.
- (backtracking) remove the last position `{i, j}` from `B`.

**References:**

- https://en.wikipedia.org/wiki/Eight_queens_puzzle#Exercise_in_algorithm_design
- https://leetcode.com/problems/n-queens/discuss/2107948/C%2B%2B-or-Easy-Explanation-w-Backtracking
