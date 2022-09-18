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

### Level Order with Queue + Reverse

比较简单直观的方法，借助 queue 进行层序遍历，根据 zigzag 的条件，反转特定一层的结果。

### Level Order with Two Stacks

借助两个 stack 达到 zigzag level order 的效果：

```shell
depth                                  stack                                                        result
 0              3                      pop st2(3), push node.left, node.right -> st1                [3]
            /       \
 1         9         8                 pop st1(9, 8), push node.right, node.left -> st2             [8, 9]
         /   \     /   \
 2      1     3   7     6              pop st2(6, 7, 3, 1), push node.left node.right -> st1        [1, 3, 7, 6]
             / \       / \
 3          4   5     2   3            pop st1(4, 5, 2, 3), push node.right, node.left -> st2       [3, 2, 5, 4]
```
