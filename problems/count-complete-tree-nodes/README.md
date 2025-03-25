# 222. Count Complete Tree Nodes

- Difficulty: Medium
- Topics: Binary Search, Tree, Depth-First Search, Binary Tree
- Link: https://leetcode.com/problems/count-complete-tree-nodes/

## Description

Given the `root` of a **complete** binary tree, return the number of the nodes in the tree.

According to **[Wikipedia](http://en.wikipedia.org/wiki/Binary_tree#Types_of_binary_trees)**, every level, except possibly the last, is completely filled in a complete binary tree, and all nodes in the last level are as far left as possible. It can have between `1` and `2h` nodes inclusive at the last level `h`.

Design an algorithm that runs in less than `O(n)` time complexity.

**Example 1:**

![img](https://assets.leetcode.com/uploads/2021/01/14/complete.jpg)

```
Input: root = [1,2,3,4,5,6]
Output: 6
```

**Example 2:**

```
Input: root = []
Output: 0
```

**Example 3:**

```
Input: root = [1]
Output: 1
```

**Constraints:**

- The number of nodes in the tree is in the range `[0, 5 * 10^4]`.
- `0 <= Node.val <= 5 * 10^4`
- The tree is guaranteed to be **complete**.

## Solution

### Approach: Binary Search on a Complete Binary Tree

- A complete binary tree has all levels fully filled except possibly the last, which is left-aligned.
- Compute the left and right subtree heights:
  - If they are equal, the left subtree is a perfect tree of height `h`, contributing `2^h - 1` nodes. Recur on the right subtree.
  - Otherwise, the right subtree is a perfect tree of height `h-1`, contributing `2^(h-1) - 1` nodes. Recur on the left subtree.
- This reduces the problem to `O(log n * log n)`, as height calculation takes `O(log n)`, and recursion runs `O(log n)` times.

**Complexity**:

- Time: `O(log n * log n)`
- Space: `O(log n)` (recursive call stack).
