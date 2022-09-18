# 938. Range Sum of BST

- Difficulty: Easy
- Topics: Tree, Depth-First Search, Binary Search Tree, Binary Tree
- Link: https://leetcode.com/problems/range-sum-of-bst/

## Description

Given the `root` node of a binary search tree and two integers `low` and `high`, return _the sum of values of all nodes with a value in the **inclusive** range_ `[low, high]`.

**Example 1:**

![img](https://assets.leetcode.com/uploads/2020/11/05/bst1.jpg)

```
Input: root = [10,5,15,3,7,null,18], low = 7, high = 15
Output: 32
Explanation: Nodes 7, 10, and 15 are in the range [7, 15]. 7 + 10 + 15 = 32.
```

**Example 2:**

![img](https://assets.leetcode.com/uploads/2020/11/05/bst2.jpg)

```
Input: root = [10,5,15,3,7,13,18,1,null,6], low = 6, high = 10
Output: 23
Explanation: Nodes 6, 7, and 10 are in the range [6, 10]. 6 + 7 + 10 = 23.
```

**Constraints:**

- The number of nodes in the tree is in the range `[1, 2 * 10^4]`.
- `1 <= Node.val <= 10^5`
- `1 <= low <= high <= 10^5`
- All `Node.val` are **unique**.

## Solution

Refer: @archit91 https://leetcode.com/problems/range-sum-of-bst/discuss/1628229

### BFS, Optimized for BST

To take advantage of the the fact that our given tree is a binary search tree and all node's value are unique, we can reduce search space in some cases by doing conditional tabulation(also called `pruning`).

- if current node's value is less than `low`, then it's useless to push the left child into the queue.
- if current node's value is greater than `high`, then it's useless to push the right child into the queue.
- So, we will only push a left or right child into the queue if current node's value is `>low` or `<high`.

### DFS, Optimized for BST

Similar to above solution. When using DFS, we can reduce search space in some cases by doing conditional recursion.
