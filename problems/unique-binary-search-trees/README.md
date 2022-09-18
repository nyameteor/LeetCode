# 96. Unique Binary Search Trees

- Difficulty: Medium
- Topics: Math, Dynamic Programming, Tree, Binary Search Tree, Binary Tree
- Link: https://leetcode.com/problems/unique-binary-search-trees/

## Description

Given an integer `n`, return _the number of structurally unique **BST'**s (binary search trees) which has exactly_ `n` _nodes of unique values from_ `1` _to_ `n`.

**Example 1:**

![img](https://assets.leetcode.com/uploads/2021/01/18/uniquebstn3.jpg)

```
Input: n = 3
Output: 5
```

**Example 2:**

```
Input: n = 1
Output: 1
```

**Constraints:**

- `1 <= n <= 19`

## Solution

### Math + Dynamic Programming

设 $A(n)$ 为给定 ${1, 2, ... n}$ 的 $n$ 个节点时 Unique BST 的数量，其中 $A(0) = 1$（类似 $0! = 1$）。

设 $C(i, n)$ 为给定 ${1, 2, ...n}$ 的 $n$ 个节点且根节点为 $i$ 时 Unique BST 的数量，按定义可知：

$$
A(n) = \sum_{i=1}^{n}C(i, n) \quad n \geq 1
$$

通过画图观察，可以得到以下规律：

当 i 为根节点时，剩余 n - 1 个节点中，小于 i 的数量为 i - 1，大于 i 的数量为 n - i，此时的组合共有 $A(i-1) * A(n-i)$ 种，即：

$$
C(i, n) = A(i-1) * A(n-i)
$$

如 n = 5，i = 3 时，小于 i 的数量为 2，大于 i 的数量为 2，$C(3, 5) = A(2) * A(2) = 2 * 2 = 4$。

由此可知，该问题有重叠的子问题，适合使用动态规划。总的方程可表示为：

$$
A(n) =
    \begin{cases}
        1 & \quad n = 0 \\
        \sum_{i=1}^{n}A(i-1)*A(n-i) & \quad n \geq 1
    \end{cases}
$$

推出递推方程后用 memoization 或 tabulation 都很简单。
