# 515. Find Largest Value in Each Tree Row

- Difficulty: Medium
- Topics: Tree, Depth-First Search, Breadth-First Search, Binary Tree
- Link: https://leetcode.com/problems/find-largest-value-in-each-tree-row/

## Description

Given the `root` of a binary tree, return _an array of the largest value in each row_ of the tree **(0-indexed)**.

**Example 1:**

![](https://assets.leetcode.com/uploads/2020/08/21/largest_e1.jpg)

```
Input: root = [1,3,2,5,3,null,9]
Output: [1,3,9]
```

**Example 2:**

```
Input: root = [1,2,3]
Output: [1,3]
```

**Constraints:**

- The number of nodes in the tree will be in the range `[0, 10^4]`.
- `-2^31 <= Node.val <= 2^31 - 1`

## Solution

Searching the tree using Breadth-first search, and finding the largest value in each level.
