# 113. Path Sum II

- Difficulty: Medium
- Topics: Backtracking, Tree, Depth-First Search, Binary Tree
- Link: https://leetcode.com/problems/path-sum-ii/

## Description

Given the `root` of a binary tree and an integer `targetSum`, return _all **root-to-leaf** paths where the sum of the node values in the path equals_ `targetSum`_. Each path should be returned as a list of the node **values**, not node references_.

A **root-to-leaf** path is a path starting from the root and ending at any leaf node. A **leaf** is a node with no children.

**Example 1:**

![img](https://assets.leetcode.com/uploads/2021/01/18/pathsumii1.jpg)

```
Input: root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
Output: [[5,4,11,2],[5,8,4,5]]
Explanation: There are two paths whose sum equals targetSum:
5 + 4 + 11 + 2 = 22
5 + 8 + 4 + 5 = 22
```

**Example 2:**

![img](https://assets.leetcode.com/uploads/2021/01/18/pathsum2.jpg)

```
Input: root = [1,2,3], targetSum = 5
Output: []
```

**Example 3:**

```
Input: root = [1,2], targetSum = 0
Output: []
```

**Constraints:**

- The number of nodes in the tree is in the range `[0, 5000]`.
- `-1000 <= Node.val <= 1000`
- `-1000 <= targetSum <= 1000`

## Solution

### Approach: Backtracking

We use **Backtracking** and **Depth-First Search (DFS)** to explore all root-to-leaf paths and find those that sum to `targetSum`.

1. **Recursive DFS**: Traverse the tree, keeping track of the current path and remaining sum.
2. **Valid Path Check**: If a leaf node is reached and the remaining sum is `0`, store the path.
3. **Backtracking**: After exploring a node’s children, remove it from the path to backtrack.
