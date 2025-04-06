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

### Observations

1. Each node has 0, 1, or 2 children, and both `p` and `q` must exist in the tree.
2. We want to search for p and q starting from the root and bubble up information about their location.

### Steps

For each node, we want to check: "Does the left subtree or right subtree contain p or q?"

- Base Case:
  - If the current node is nil, return nil.
  - If the current node is `p` or `q`, return the current node — it might be the ancestor.

- Recursive Case:
  - If both left and right return non-nil, this means:
    - `p` was found in one subtree.
    - `q` was found in the other.
    - So the current node is the lowest common ancestor.
  - If only one side is non-nil, bubble that node up.
  - If both are nil, bubble nil up.
