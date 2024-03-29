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

给定一个二叉树，检查它是否是围绕其中心镜像对称的。

这题和 [100. Same Tree](https://leetcode.com/problems/same-tree/) 可以用相似的思路。

### Recursion

二叉树是否以中心为镜像对称，可以通过检查 root 的左右子树是否对称来判断：

```shell
      2         2
     / \       / \
    3   4     4   3
   / \           / \
  5   6         6   5
 subtree-p    subtree-q
```

很明显，递归的基本情况是：

```shell
p, q 都为 null, 返回 true             (p, q 相等，对称)
p, q 中有且只有一个为 null, 返回 false  (p, q 不等，不对称)
```

当 p, q 不是基本情况时，递推关系为：

```shell
p 与 q 是否对称 = p 和 q 的值相等 and
                p 的左子树与 q 的右子树对称 and
                p 的右子树与 q 的左子树对称
```

### Iteration

Todo
