# 101. Symmetric Tree

- Difficulty: Easy
- Topics: Tree, Depth-First Search, Breadth-First Search, Binary Tree
- Link: https://leetcode.com/problems/symmetric-tree/

## Description

Given the `root` of a binary tree, _check whether it is a mirror of itself_ (i.e., symmetric around its center).

**Example 1:**

![img](https://assets.leetcode.com/uploads/2021/02/19/symtree1.jpg)

```
Input: root = [1,2,2,3,4,4,3]
Output: true
```

**Example 2:**

![img](https://assets.leetcode.com/uploads/2021/02/19/symtree2.jpg)

```
Input: root = [1,2,2,null,3,null,3]
Output: false
```

**Constraints:**

- The number of nodes in the tree is in the range `[1, 1000]`.
- `-100 <= Node.val <= 100`

**Follow up:** Could you solve it both recursively and iteratively?

## Solution

### Approach: Recursion

We can use **recursion** to check if the left and right subtrees are mirror images of each other.

- If both subtrees are `nil`, they are symmetric.
- If only one is `nil`, they are not symmetric.
- Otherwise, their values must be equal, and their corresponding left-right and right-left child pairs must also be symmetric.

This ensures the tree is a mirror of itself.
