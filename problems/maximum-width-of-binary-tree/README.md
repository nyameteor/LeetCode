# 662. Maximum Width of Binary Tree

- Difficulty: Medium
- Topics: Tree, Depth-First Search, Breadth-First Search, Binary Tree
- Link: https://leetcode.com/problems/maximum-width-of-binary-tree/

## Description

Given the `root` of a binary tree, return _the **maximum width** of the given tree_.

The **maximum width** of a tree is the maximum **width** among all levels.

The **width** of one level is defined as the length between the end-nodes (the leftmost and rightmost non-null nodes), where the null nodes between the end-nodes are also counted into the length calculation.

It is **guaranteed** that the answer will in the range of **32-bit** signed integer.

**Example 1:**

![img](https://assets.leetcode.com/uploads/2021/05/03/width1-tree.jpg)

```
Input: root = [1,3,2,5,3,null,9]
Output: 4
Explanation: The maximum width existing in the third level with the length 4 (5,3,null,9).
```

**Example 2:**

![img](https://assets.leetcode.com/uploads/2021/05/03/width2-tree.jpg)

```
Input: root = [1,3,null,5,3]
Output: 2
Explanation: The maximum width existing in the third level with the length 2 (5,3).
```

**Example 3:**

![img](https://assets.leetcode.com/uploads/2021/05/03/width3-tree.jpg)

```
Input: root = [1,3,2,5]
Output: 2
Explanation: The maximum width existing in the second level with the length 2 (3,2).
```

**Constraints:**

- The number of nodes in the tree is in the range `[1, 3000]`.
- `-100 <= Node.val <= 100`

## Solution

The idea is simple, as shown below, the number is the `index` of node.

```shell
                                1
            +-------------------+-------------------+
            2                                       3
     +------+------+                         +------+------+
     4             5                         6             7
  +--+--+       +--+--+                   +--+--+       +--+--+
  8     9       10    11                  12    13      14    15
```

From observation, regardless whether these nodes exist:

- the index of left child is `current_index * 2`
- the index of right child is `current_index * 2 + 1`

The base case is the root, the index of root is `1`.

So we can just:

- Traversal the nodes by level, and **record** the index of child of the current node.
- In each level, we find the `width` between leftmost and rightmost index, and compare it with current `max_width`.

Refer:

- https://leetcode.com/problems/maximum-width-of-binary-tree/discuss/106645/C%2B%2BJava-*-BFSDFS3liner-Clean-Code-With-Explanation
- https://leetcode.com/problems/maximum-width-of-binary-tree/discuss/1338308/C%2B%2B-DFS-and-BFS-with-no-overflow

Similar Problem:

- [655. Print Binary Tree.md](./655.%20Print%20Binary%20Tree.md)
