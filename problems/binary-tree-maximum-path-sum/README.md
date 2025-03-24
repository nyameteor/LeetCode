# 124. Binary Tree Maximum Path Sum

- Difficulty: Hard
- Topics: Dynamic Programming, Tree, Depth-First Search, Binary Tree
- Link: https://leetcode.com/problems/binary-tree-maximum-path-sum/

## Description

A **path**
in a binary tree is a sequence of nodes where each pair of adjacent
nodes in the sequence has an edge connecting them. A node can only
appear in the sequence **at most once**. Note that the path does not need to pass through the root.

The **path sum** of a path is the sum of the node's values in the path.

Given the `root` of a binary tree, return *the maximum **path sum** of any **non-empty** path*.

**Example 1:**

![](https://assets.leetcode.com/uploads/2020/10/13/exx1.jpg)

```
Input: root = [1,2,3]
Output: 6
Explanation: The optimal path is 2 -> 1 -> 3 with a path sum of 2 + 1 + 3 = 6.
```

**Example 2:**

![](https://assets.leetcode.com/uploads/2020/10/13/exx2.jpg)

```
Input: root = [-10,9,20,null,null,15,7]
Output: 42
Explanation: The optimal path is 15 -> 20 -> 7 with a path sum of 15 + 20 + 7 = 42.
```

**Constraints:**

- The number of nodes in the tree is in the range `[1, 3 * 10^4]`.
- `-1000 <= Node.val <= 1000`

## Solution

### Approach: Depth-First Search

We use **DFS** to compute the maximum path sum:

- At each node, compute the **max gain** from left and right children, ignoring negative sums.
- Update `maxSum` with the path sum using the current node as the root (`root.Val + left + right`).
- Return the **max gain** from one side (`root.Val + max(left, right)`) to continue the path.

**Complexity**:

- Time Complexity: `O(n)` (each node visited once).
- Space Complexity: `O(h)` (recursion depth, worst case `O(n)` for a skewed tree).
