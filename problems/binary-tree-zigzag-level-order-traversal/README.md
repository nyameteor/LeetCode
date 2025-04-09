# 103. Binary Tree Zigzag Level Order Traversal

- Difficulty: Medium
- Topics: Tree, Breadth-First Search, Binary Tree
- Link: https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/

## Description

Given the `root` of a binary tree, return _the zigzag level order traversal of its nodes' values_. (i.e., from left to right, then right to left for the next level and alternate between).

**Example 1:**

![img](https://assets.leetcode.com/uploads/2021/02/19/tree1.jpg)

```
Input: root = [3,9,20,null,null,15,7]
Output: [[3],[20,9],[15,7]]
```

**Example 2:**

```
Input: root = [1]
Output: [[1]]
```

**Example 3:**

```
Input: root = []
Output: []
```

**Constraints:**

- The number of nodes in the tree is in the range `[0, 2000]`.
- `-100 <= Node.val <= 100`

## Solution

- Use a queue (BFS) to process nodes level by level.
- A flag (`even`) tracks whether to reverse the order for each level.
- For each level, collect node values and reverse if needed, then add them to the result.
