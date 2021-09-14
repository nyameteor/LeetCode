# 236. Lowest Common Ancestor of a Binary Tree

- Difficulty: Medium
- Topics: Tree, Depth-First Search, Binary Tree
- Link: https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/

## Description

Given a binary tree, find the lowest common ancestor (LCA) of two given nodes in the tree.

According to the [definition of LCA on Wikipedia](https://en.wikipedia.org/wiki/Lowest_common_ancestor): “The lowest common ancestor is defined between two nodes `p` and `q` as the lowest node in `T` that has both `p` and `q` as descendants (where we allow **a node to be a descendant of itself**).”

**Example 1:**

![img](https://assets.leetcode.com/uploads/2018/12/14/binarytree.png)

```
Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
Output: 3
Explanation: The LCA of nodes 5 and 1 is 3.
```

**Example 2:**

![img](https://assets.leetcode.com/uploads/2018/12/14/binarytree.png)

```
Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
Output: 5
Explanation: The LCA of nodes 5 and 4 is 5, since a node can be a descendant of itself according to the LCA definition.
```

**Example 3:**

```
Input: root = [1,2], p = 1, q = 2
Output: 1
```

**Constraints:**

- The number of nodes in the tree is in the range `[2, 10^5]`.
- `-10^9 <= Node.val <= 10^9`
- All `Node.val` are **unique**.
- `p != q`
- `p` and `q` will exist in the tree.

## Solution

给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

### Recursion

按照*最近公共祖先*的定义，可发现：

- 找到 p, q 节点的路径 path_p, path_q
- p, q 的最近公共祖先即是 path_p 和 path_q 中最后一个相同的节点。

```shell
# Example 1
path_p = [3, 5]
path_q = [3, 1]
lca = 3

# Example 2
path_p = [3, 5]
path_q = [3, 5, 2, 4]
lca = 5
```

故可以创建一个 `findPath` 方法，获取 p 和 q 的 path 后遍历返回其最后一个相同的节点。

Todo: findPath 方法说明
