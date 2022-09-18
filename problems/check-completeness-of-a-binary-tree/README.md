# 958. Check Completeness of a Binary Tree

- Difficulty: Medium
- Topics: Tree, Breadth-First Search, Binary Tree
- Link: https://leetcode.com/problems/check-completeness-of-a-binary-tree/

## Description

Given the `root` of a binary tree, determine if it is a _complete binary tree_.

In a **[complete binary tree](http://en.wikipedia.org/wiki/Binary_tree#Types_of_binary_trees)**, every level, except possibly the last, is completely filled, and all nodes in the last level are as far left as possible. It can have between `1` and `2h` nodes inclusive at the last level `h`.

**Example 1:**

![img](https://assets.leetcode.com/uploads/2018/12/15/complete-binary-tree-1.png)

```
Input: root = [1,2,3,4,5,6]
Output: true
Explanation: Every level before the last is full (ie. levels with node-values {1} and {2, 3}), and all nodes in the last level ({4, 5, 6}) are as far left as possible.
```

**Example 2:**

![img](https://assets.leetcode.com/uploads/2018/12/15/complete-binary-tree-2.png)

```
Input: root = [1,2,3,4,5,null,7]
Output: false
Explanation: The node with value 7 isn't as far left as possible.
```

**Constraints:**

- The number of nodes in the tree is in the range `[1, 100]`.
- `1 <= Node.val <= 1000`

## Solution

判断一棵二叉树的完整性。在一颗二叉树中，若除最后一层外的其余层都是满的，并且最后一层要么是满的，要么在右边缺少连续若干节点，则此二叉树为完全二叉树（Complete Binary Tree）。

判断的过程思考得不清楚，需要再看。

- 层序遍历，若节点为 null，则设二叉树的 isEnd 为 true
- 若下次循环时 isEnd 为 true，则说明有多个 End，不满足完全二叉树定义，返回 false

Refer:

- [Wikipedia/Types of binary trees](https://en.wikipedia.org/wiki/Binary_tree#Types_of_binary_trees)
- [leetcode/dicuss/1437393](https://leetcode.com/problems/check-completeness-of-a-binary-tree/discuss/1437393/C%2B%2B-bfs)
