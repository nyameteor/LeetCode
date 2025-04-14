# 96. Unique Binary Search Trees

- Difficulty: Medium
- Topics: Math, Dynamic Programming, Tree, Binary Search Tree, Binary Tree
- Link: https://leetcode.com/problems/unique-binary-search-trees/

## Description

Given an integer `n`, return _the number of structurally unique **BST'**s (binary search trees) which has exactly_ `n` _nodes of unique values from_ `1` _to_ `n`.

**Example 1:**

![img](https://assets.leetcode.com/uploads/2021/01/18/uniquebstn3.jpg)

```
Input: n = 3
Output: 5
```

**Example 2:**

```
Input: n = 1
Output: 1
```

**Constraints:**

- `1 <= n <= 19`

## Solution

### Key Idea

This problem is a classic [Catalan number](https://en.wikipedia.org/wiki/Catalan_number) application.

To build a BST with `n` nodes:

- Choose each node as the root.
- Recursively count valid left/right subtrees (`left + right = n - 1`).

Recurrence relation:

```
numTrees(n) = sum ((numTrees(left) * numTrees(right)) for all left in [0, n-1])
```

Use **top-down DP with memoization** to cache results and avoid recomputation.
