# 104. Maximum Depth of Binary Tree

- Difficulty: Easy
- Topics: Tree, Depth-First Search, Breadth-First Search, Binary Tree
- Link: https://leetcode.com/problems/maximum-depth-of-binary-tree/

## Description

Given the `root` of a binary tree, return _its maximum depth_.

A binary tree's **maximum depth** is the number of nodes along the longest path from the root node down to the farthest leaf node.

**Example 1:**

![img](https://assets.leetcode.com/uploads/2020/11/26/tmp-tree.jpg)

```
Input: root = [3,9,20,null,null,15,7]
Output: 3
```

**Example 2:**

```
Input: root = [1,null,2]
Output: 2
```

**Constraints:**

- The number of nodes in the tree is in the range `[0, 104]`.
- `-100 <= Node.val <= 100`

## Solution

### Depth-First Search

二叉树的最大深度 = 左右子树的最大深度 + 当前深度，若节点为空，则深度为 0。可使用深度优先搜索。

### Breadth-First Search

二叉树的最大深度为广度优先搜索的层数。