# 404. Sum of Left Leaves

- Difficulty: Easy
- Topics: Tree, Depth-First Search, Breadth-First Search, Binary Tree
- Link: https://leetcode.com/problems/sum-of-left-leaves/

## Description

Given the `root` of a binary tree, return the sum of all left leaves.

**Example 1:**

![img](https://assets.leetcode.com/uploads/2021/04/08/leftsum-tree.jpg)

```
Input: root = [3,9,20,null,null,15,7]
Output: 24
Explanation: There are two left leaves in the binary tree, with values 9 and 15 respectively.
```

**Example 2:**

```
Input: root = [1]
Output: 0
```

**Constraints:**

- The number of nodes in the tree is in the range `[1, 1000]`.
- `-1000 <= Node.val <= 1000`

## Solution

Use DFS to traverse the tree and track whether the current node is a left child. If it's a leaf node and marked as left, add its value to the sum.
