# 226. Invert Binary Tree

- Difficulty: Easy
- Topics: Tree, Depth-First Search, Breadth-First Search, Binary Tree
- Link: https://leetcode.com/problems/invert-binary-tree/

## Description

Given the `root` of a binary tree, invert the tree, and return _its root_.

**Example 1:**

![img](https://assets.leetcode.com/uploads/2021/03/14/invert1-tree.jpg)

```
Input: root = [4,2,7,1,3,6,9]
Output: [4,7,2,9,6,3,1]
```

**Example 2:**

![img](https://assets.leetcode.com/uploads/2021/03/14/invert2-tree.jpg)

```
Input: root = [2,1,3]
Output: [2,3,1]
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

### Approach: Recursion

- **Swap** the left and right children of each node.
- **Recursively** invert the left and right subtrees.
- **Base case**: Return `nil` if the node is `nil`.

**Time Complexity**: **O(n)**, where `n` is the number of nodes.

**Space Complexity**: **O(h)**, where `h` is the height of the tree (due to recursion stack).
