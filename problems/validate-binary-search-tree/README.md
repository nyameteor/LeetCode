# 98. Validate Binary Search Tree

- Difficulty: Medium
- Topics: Tree, Depth-First Search, Binary Search Tree, Binary Tree
- Link: https://leetcode.com/problems/validate-binary-search-tree/

## Description

Given the `root` of a binary tree, _determine if it is a valid binary search tree (BST)_.

A **valid BST** is defined as follows:

- The left subtree of a node contains only nodes with keys **less than** the node's key.
- The right subtree of a node contains only nodes with keys **greater than** the node's key.
- Both the left and right subtrees must also be binary search trees.

**Example 1:**

![img](https://assets.leetcode.com/uploads/2020/12/01/tree1.jpg)

```
Input: root = [2,1,3]
Output: true
```

**Example 2:**

![img](https://assets.leetcode.com/uploads/2020/12/01/tree2.jpg)

```
Input: root = [5,1,4,null,null,3,6]
Output: false
Explanation: The root node's value is 5 but its right child's value is 4.
```

**Constraints:**

- The number of nodes in the tree is in the range `[1, 10^4]`.
- `-2^31 <= Node.val <= 2^31 - 1`

## Solution

### Observations

In a **Binary Search Tree (BST)**, each node must satisfy:

- All nodes in the left subtree are **strictly less** than the current node.
- All nodes in the right subtree are **strictly greater** than the current node.

### Approach

- Perform DFS with value range constraints.
- At each node, check: `minVal < node.Val < maxVal`.
- Recursively update bounds:
  - Left child: `maxVal = node.Val`
  - Right child: `minVal = node.Val`

This enforces global BST rules, not just local relationships.
