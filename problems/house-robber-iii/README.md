# 337. House Robber III

- Difficulty: Medium
- Topics: Dynamic Programming, Tree, Depth-First Search, Binary Tree
- Link: https://leetcode.com/problems/house-robber-iii/

## Description

The thief has found himself a new place for his thievery again. There is only one entrance to this area, called `root`.

Besides the `root`, each house has one and only one parent house. After a tour, the smart thief realized that all houses in this place form a binary tree. It will automatically contact the police if **two directly-linked houses were broken into on the same night**.

Given the `root` of the binary tree, return \*the maximum amount of money the thief can rob **without alerting the police\***.

**Example 1:**

![img](https://assets.leetcode.com/uploads/2021/03/10/rob1-tree.jpg)

```
Input: root = [3,2,3,null,3,null,1]
Output: 7
Explanation: Maximum amount of money the thief can rob = 3 + 3 + 1 = 7.
```

**Example 2:**

![img](https://assets.leetcode.com/uploads/2021/03/10/rob2-tree.jpg)

```
Input: root = [3,4,5,1,3,null,1]
Output: 9
Explanation: Maximum amount of money the thief can rob = 4 + 5 = 9.
```

**Constraints:**

- The number of nodes in the tree is in the range `[1, 10^4]`.
- `0 <= Node.val <= 10^4`

## Solution

### Recursion + Memoization

设偷盗以 $A$ 为根结点的树能获益的最大值为 $C(A)$。

$C(A)$ 的方程：

$$
C(A) =
    \begin{cases}
        0 & A = null  \\
        A.val & A.left = null \land A.right = null\\
        max\{ C(A.left) + C(A.right),
            A.val + (C(A.left.left) + C(A.left.right))
            + (C(A.right.left) + C(A.right.right))
        \}  & A.left \lor A.right
    \end{cases}
$$
