# 226. Invert Binary Tree

- Difficulty: Easy
- Topics: Tree, Depth-First Search, Breadth-First Search, Binary Tree
- Link: https://leetcode.com/problems/invert-binary-tree/

## Description

Given the `root` of a binary tree, invert the tree, and return _its root_.

**Example 1:**

![img](https://assets.leetcode.com/uploads/2021/03/14/invert1-tree.jpg)

```
Input: root = [4,2,7,1,3,6,9]
Output: [4,7,2,9,6,3,1]
```

**Example 2:**

![img](https://assets.leetcode.com/uploads/2021/03/14/invert2-tree.jpg)

```
Input: root = [2,1,3]
Output: [2,3,1]
```

**Example 3:**

```
Input: root = []
Output: []
```

**Constraints:**

- The number of nodes in the tree is in the range `[0, 100]`.
- `-100 <= Node.val <= 100`

## Solution

### Recursion

- 若根节点为 null，返回 null
- 若根节点不为 null
  1. 递归调用反转根节点的左孩子
  2. 递归调用反转根节点的右孩子
  3. 左右交换根节点的左孩子和右孩子

### Iteration

Todo
