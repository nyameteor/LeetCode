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

按照 Binary Search Tree 的定义：

- 左子树**所有**节点的 val 均小于 node 的 val
- 右子树**所有**节点的 val 均大于 node 的 val
- 左右子树也是 Binary Search Tree

### Depth-First Search with in-order

对 Binary Search Tree 进行 in-order 遍历，可以发现序列是**从小到大有序**的，可依此来判断。
