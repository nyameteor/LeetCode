# 110. Balanced Binary Tree

- Difficulty: Easy
- Topics: Tree, Depth-First Search, Binary Tree
- Link: https://leetcode.com/problems/balanced-binary-tree/

## Description

Given a binary tree, determine if it is height-balanced.

For this problem, a height-balanced binary tree is defined as:

> a binary tree in which the left and right subtrees of _every_ node differ in height by no more than 1.

**Example 1:**

![](https://assets.leetcode.com/uploads/2020/10/06/balance_1.jpg)

```
Input: root = [3,9,20,null,null,15,7]
Output: true
```

**Example 2:**

![](https://assets.leetcode.com/uploads/2020/10/06/balance_2.jpg)

```
Input: root = [1,2,2,3,3,null,null,4,4]
Output: false
```

**Example 3:**

```
Input: root = []
Output: true
```

**Constraints:**

- The number of nodes in the tree is in the range `[0, 5000]`.
- `-10^4 <= Node.val <= 10^4`

## Solution

This problem has two approaches: top down and bottom up way.

For me, finding out the bottom up recursive way was not `easy`.

**References**

- [The bottom up O(N) solution would be better](https://leetcode.com/problems/balanced-binary-tree/solutions/35691/the-bottom-up-on-solution-would-be-better/)
- [Golang DFS solution (bottom up)](https://leetcode.com/problems/balanced-binary-tree/solutions/35732/golang-dfs-solution-bottom-up)
