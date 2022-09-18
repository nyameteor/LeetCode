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

### Recursive DFS

本题给定了一棵二叉树，求在所有为左孩子的叶子的和。

判断叶子节点很简单，那么如何判定当前节点是左孩子还是右孩子呢？一个实用的方法是，在遍历时添加一个 bool 变量，用于跟踪当前节点是左孩子还是右孩子：

- 从节点 root 开始，初始化 bool 变量 `isLeft=false`，因为 root 节点不是某个节点的左孩子。
- 递归探索左子树时，设置 `isLeft=true`；递归探索右子树时，设置 `isLeft=false`。
- 若当前节点是叶子节点且 `isLeft=true`，返回当前节点的值；否则，返回 0。

Refer: https://leetcode.com/problems/sum-of-left-leaves/discuss/1558055
