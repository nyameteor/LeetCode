# 199. Binary Tree Right Side View

- Difficulty: Medium
- Topics: Tree, Depth-First Search, Breadth-First Search, Binary Search
- Link: https://leetcode.com/problems/binary-tree-right-side-view/

## Description

Given the `root` of a binary tree, imagine yourself standing on the **right side** of it, return _the values of the nodes you can see ordered from top to bottom_.

**Example 1:**

![img](https://assets.leetcode.com/uploads/2021/02/14/tree.jpg)

```
Input: root = [1,2,3,null,5,null,4]
Output: [1,3,4]
```

**Example 2:**

```
Input: root = [1,null,3]
Output: [1,3]
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

二叉树的右视图，即返回的结果序列 res 为二叉树每一层最靠右的节点。

### Depth-First Search

DFS 时添加 level 变量记录当前层数（深度）。

递归时先右子树，再左子树。将当前 level 中访问的第一个节点添加到 res 中。

### Breadth-First Search

BFS 时添加 level 变量记录当前层数（深度），并添加一个 levelSize 变量用于辅助计算 level。

levelSize 记为当前 queue 的大小（即当前 level 节点的数量），循环 levelSize 次后本层的节点全部出队，下层的节点全部入队，此时 level++。

遍历时先右子树，再左子树。将当前 level 中访问的第一个节点添加到 res 中。
