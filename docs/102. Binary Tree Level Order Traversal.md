# 102. Binary Tree Level Order Traversal

- Difficulty: Medium
- Topics: Tree, Breadth-First Search, Binary Tree
- Link: https://leetcode.com/problems/binary-tree-level-order-traversal/

## Description

Given the `root` of a binary tree, return _the level order traversal of its nodes' values_. (i.e., from left to right, level by level).

**Example 1:**

![img](https://assets.leetcode.com/uploads/2021/02/19/tree1.jpg)

```
Input: root = [3,9,20,null,null,15,7]
Output: [[3],[9,20],[15,7]]
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
- `-1000 <= Node.val <= 1000`

## Solution

给一个二叉树，请返回其按广度优先遍历得到的节点值（即逐层地，从左到右访问所有节点）。

这道题可以作为二叉树广度优先遍历的板子，以此为基础解决许多问题。

### Iteration

- 深度优先遍历可以用 `stack` 作为辅助数据结构，即符合递归的逻辑
- 广度优先遍历可以用 `queue` 作为辅助数据结构，即符合逐层遍历的逻辑

参照 Wikipedia [Breadth-first_search](https://en.wikipedia.org/wiki/Tree_traversal#Breadth-first_search) 伪代码，由于本题要返回每层的 node，需在其中再嵌套一个循环将结果分层：

```shell
levelorder(root)
    q ← empty queue
    q.enqueue(root)
    while not q.isEmpty() do
        len ← q.size()
        while len > 0 do
          node ← q.dequeue()
          visit(node)
          if node.left ≠ null then
              q.enqueue(node.left)
          if node.right ≠ null then
              q.enqueue(node.right)
          len ← len - 1
```
