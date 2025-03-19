# 199. Binary Tree Right Side View

- Difficulty: Medium
- Topics: Tree, Depth-First Search, Breadth-First Search, Binary Search
- Link: https://leetcode.com/problems/binary-tree-right-side-view/

## Description

Given the `root` of a binary tree, imagine yourself standing on the **right side** of it, return _the values of the nodes you can see ordered from top to bottom_.

**Example 1:**

![img](https://assets.leetcode.com/uploads/2021/02/14/tree.jpg)

```
Input: root = [1,2,3,null,5,null,4]
Output: [1,3,4]
```

**Example 2:**

```
Input: root = [1,null,3]
Output: [1,3]
```

**Example 3:**

```
Input: root = []
Output: []
```

**Constraints:**

- The number of nodes in the tree is in the range `[0, 100]`.
- `-100 <= Node.val <= 100`

## Solution

### Approach: Depth-First Search

1. Traverse the **right subtree first** to prioritize rightmost nodes at each level.
2. Append the first node encountered at each level to the result list.
3. Use a helper function to track depth and update the result.
