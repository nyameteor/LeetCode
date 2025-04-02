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

This problem can be solved using **Dynamic Programming** on a binary tree. For each node, there are two options:

1. **Rob the node**: Skip its children and add their "not robbed" values.
2. **Don't rob the node**: Consider the maximum of robbing or not robbing its children.

**Steps:**

- Use a postorder DFS traversal to compute, at each node, the maximum money from both choices.
- Propagate the results upwards and return the maximum value for the root.

**Complexity:**

- **Time:** O(n), where n is the number of nodes.
- **Space:** O(h), where h is the tree height due to recursion.

Alternatively, a **top-down DP with memoization** approach can be used, but it introduces additional overhead due to the hashmap.
