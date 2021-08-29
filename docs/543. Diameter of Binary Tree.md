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

给定二叉树的根，返回树的直径长度。二叉树的直径是树中任意两个节点之间最长路径的长度。此路径可能会也可能不会通过根。

### Recursion

看到这个问题画图，首先想到的是二叉树的直径 = 左右子树的深度之和，容易得出：

```shell
# 基本情况
node == null, diameter(node) = 0

# 递推方程
node != null, diameter(node) = 1 + max(diameter(node.left), diameter(node.right))
```

但这样只符合路径经过根节点的情况，实际路径不一定经过根节点，如在这个情况，[8, 6, 4, 2, 5, 7, 9] 的路径长度大于 [8, 6, 4, 2, 1, 3]：

```shell
        1
       / \
      2   3
     / \
    4   5
   /     \
  6       7
 /         \
8           9
```

故可以设置一个额外的全局变量 maximum，递归时更新 `maximum = max(maximum, 左右子树深度之和)`，最后答案返回 maximum。
