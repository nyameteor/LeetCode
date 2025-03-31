# 95. Unique Binary Search Trees II

- Difficulty: Medium
- Topics: Dynamic Programming, Backtracking, Tree, Binary Search Tree, Binary Tree
- Link: https://leetcode.com/problems/unique-binary-search-trees-ii/

## Description

Given an integer `n`, return *all the structurally unique **BST'**s (binary search trees), which has exactly* `n` *nodes of unique values from* `1` *to* `n`. Return the answer in **any order**.

**Example 1:**

![](https://assets.leetcode.com/uploads/2021/01/18/uniquebstn3.jpg)

```
Input: n = 3
Output: [[1,null,2,null,3],[1,null,3,2],[2,1,3],[3,1,null,null,2],[3,2,null,1]]
```

**Example 2:**

```
Input: n = 1
Output: [[1]]
```

**Constraints:**

- `1 <= n <= 8`

## Solution

### Observations

**Subtree Structure**: A BST can be constructed by selecting each number from 1 to `n` as the root once, and then recursively constructing the left and right subtrees. The left subtree will consist of numbers less than the root, and the right subtree will consist of numbers greater than the root.

### Steps

To generate all unique binary search trees (BSTs) for values from 1 to `n`, we can use a recursive approach:

1. **Root Selection**: For each number `i` from `1` to `n`, consider `i` as the root of the tree.
2. **Left and Right Subtrees**:
   - Left subtree: all BSTs formed by numbers in the range `[low, i-1]`.
   - Right subtree: all BSTs formed by numbers in the range `[i+1, high]`.
3. **Recursive Combination**: Combine each left and right subtree with `i` as the root to form a tree.
4. **Base Case**: When `low > high`, return `nil`, representing an empty subtree.

This approach ensures all structurally unique BSTs are generated for a given `n`.
