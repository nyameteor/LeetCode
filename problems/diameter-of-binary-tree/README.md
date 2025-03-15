# 543. Diameter of Binary Tree

- Difficulty: Easy
- Topics: Tree, Depth-First Search, Binary Tree
- Link: https://leetcode.com/problems/diameter-of-binary-tree/

## Description

Given the `root` of a binary tree, return _the length of the **diameter** of the tree_.

The **diameter** of a binary tree is the **length** of the longest path between any two nodes in a tree. This path may or may not pass through the `root`.

The **length** of a path between two nodes is represented by the number of edges between them.

**Example 1:**

![img](https://assets.leetcode.com/uploads/2021/03/06/diamtree.jpg)

```
Input: root = [1,2,3,4,5]
Output: 3
Explanation: 3 is the length of the path [4,2,1,3] or [5,2,1,3].
```

**Example 2:**

```
Input: root = [1,2]
Output: 1
```

**Constraints:**

- The number of nodes in the tree is in the range `[1, 10^4]`.
- `-100 <= Node.val <= 100`

## Solution

### Intuition

- The **diameter** of a tree is the longest path between any two nodes, which can be thought of as passing through some node’s left and right subtrees.
- For each node, the **diameter** passing through it is the sum of its **left subtree height** and **right subtree height**.
- To efficiently compute this, we use **post-order DFS** (bottom-up), where each node calculates its height while simultaneously updating the maximum diameter encountered.
