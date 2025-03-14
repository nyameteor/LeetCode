# 144. Binary Tree Preorder Traversal

- Difficulty: Easy
- Topics: Stack, Tree, Depth-First Search, Binary Tree
- Link: https://leetcode.com/problems/binary-tree-preorder-traversal/

## Description

Given the `root` of a binary tree, return _the preorder traversal of its nodes' values_.

**Example 1:**

![img](https://assets.leetcode.com/uploads/2020/09/15/inorder_1.jpg)

```
Input: root = [1,null,2,3]
Output: [1,2,3]
```

**Example 2:**

```
Input: root = []
Output: []
```

**Example 3:**

```
Input: root = [1]
Output: [1]
```

**Example 4:**

![img](https://assets.leetcode.com/uploads/2020/09/15/inorder_5.jpg)

```
Input: root = [1,2]
Output: [1,2]
```

**Example 5:**

![img](https://assets.leetcode.com/uploads/2020/09/15/inorder_4.jpg)

```
Input: root = [1,null,2]
Output: [1,2]
```

**Constraints:**

- The number of nodes in the tree is in the range `[0, 100]`.
- `-100 <= Node.val <= 100`

## Solution

References: https://en.wikipedia.org/wiki/Tree_traversal#Depth-first_search_implementation
