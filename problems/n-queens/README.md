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

### Approach: Backtracking

The N-Queens problem can be solved using **Backtracking**, which explores all potential placements of queens row by row.

1. **DFS**: Start by placing a queen in each column of the current row. Recursively try placing queens in subsequent rows.
2. **Backtracking**: If no valid placements are found, backtrack by removing the last placed queen and trying the next column.
3. **Validity Check**: A queen's position `(i, j)` is valid if:
   - No other queen shares the same column.
   - No other queen is on the same diagonal.
4. **Board Construction**: Once a valid configuration is found, construct the board and store the result.

### References

- https://en.wikipedia.org/wiki/Eight_queens_puzzle#Exercise_in_algorithm_design
- https://leetcode.com/problems/n-queens/discuss/2107948/C%2B%2B-or-Easy-Explanation-w-Backtracking
