# 222. Count Complete Tree Nodes

- Difficulty: Medium
- Topics: Binary Search, Tree, Depth-First Search, Binary Tree
- Link: https://leetcode.com/problems/count-complete-tree-nodes/

## Description

Given the `root` of a **complete** binary tree, return the number of the nodes in the tree.

According to **[Wikipedia](http://en.wikipedia.org/wiki/Binary_tree#Types_of_binary_trees)**, every level, except possibly the last, is completely filled in a complete binary tree, and all nodes in the last level are as far left as possible. It can have between `1` and `2h` nodes inclusive at the last level `h`.

Design an algorithm that runs in less than `O(n)` time complexity.

**Example 1:**

![img](https://assets.leetcode.com/uploads/2021/01/14/complete.jpg)

```
Input: root = [1,2,3,4,5,6]
Output: 6
```

**Example 2:**

```
Input: root = []
Output: 0
```

**Example 3:**

```
Input: root = [1]
Output: 1
```

**Constraints:**

- The number of nodes in the tree is in the range `[0, 5 * 10^4]`.
- `0 <= Node.val <= 5 * 10^4`
- The tree is guaranteed to be **complete**.

## Solution

### Backtracking

比较简单直观的思路：

1. 计算树的最大层数；
2. 统计最后一层的节点数：如果一个节点是叶子节点，且当前层是当前最大层，那么该节点在最后一层；
3. 节点数 = 2^(最大层数 -1) - 1 + 最后一层的节点数。

```shell
          1
      /       \
     2         3
   /   \     /   \
  4     5   6     7
 / \
8   9
```

可以加上剪枝，比如上图中深度搜索到 5 的时候，5 虽然是叶子，但所在层数小于树的最大层数，所以之后可以不继续搜索了。
