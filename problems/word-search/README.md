# 79. Word Search

- Difficulty: Medium
- Topics: Array, Backtracking, Matrix
- Link: https://leetcode.com/problems/word-search/

## Description

Given an `m x n` grid of characters `board` and a string `word`, return `true` _if_ `word` _exists in the grid_.

The word can be constructed from letters of sequentially adjacent
cells, where adjacent cells are horizontally or vertically neighboring.
The same letter cell may not be used more than once.

**Example 1:**

![](https://assets.leetcode.com/uploads/2020/11/04/word2.jpg)

```
**Input:** board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
**Output:** true
```

**Example 2:**

![](https://assets.leetcode.com/uploads/2020/11/04/word-1.jpg)

```
**Input:** board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "SEE"
**Output:** true
```

**Example 3:**

![](https://assets.leetcode.com/uploads/2020/10/15/word3.jpg)

```
**Input:** board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCB"
**Output:** false
```

**Constraints:**

- `m == board.length`
- `n = board[i].length`
- `1 <= m, n <= 6`
- `1 <= word.length <= 15`
- `board` and `word` consists of only lowercase and uppercase English letters.

**Follow up:** Could you use search pruning to make your solution faster with a larger `board`?

## Solution

Because the same cell can only be used once, we use `seen` to mark whether the cell has been visited.

We search the way along 4 directions, if current cell out of area or has been visited or not match, then this way fails.
If with current cell we can not find the way, backtracking, and search for another way, util we find a way or fails.

Reference:

https://leetcode.com/problems/word-search/solutions/649588/backtracking-runtime-4-0ms-beat-96-6/
