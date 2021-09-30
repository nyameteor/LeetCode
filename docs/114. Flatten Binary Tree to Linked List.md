# 114. Flatten Binary Tree to Linked List

- Difficulty: Medium
- Topics: Linked List, Stack, Tree, Depth-First Search, Binary Tree
- Link: https://leetcode.com/problems/flatten-binary-tree-to-linked-list/

## Description

Given the `root` of a binary tree, flatten the tree into a "linked list":

- The "linked list" should use the same `TreeNode` class where the `right` child pointer points to the next node in the list and the `left` child pointer is always `null`.
- The "linked list" should be in the same order as a [**pre-order** **traversal**](https://en.wikipedia.org/wiki/Tree_traversal#Pre-order,_NLR) of the binary tree.

**Example 1:**

![img](https://assets.leetcode.com/uploads/2021/01/14/flaten.jpg)

```
Input: root = [1,2,5,3,4,null,6]
Output: [1,null,2,null,3,null,4,null,5,null,6]
```

**Example 2:**

```
Input: root = []
Output: []
```

**Example 3:**

```
Input: root = [0]
Output: [0]
```

**Constraints:**

- The number of nodes in the tree is in the range `[0, 2000]`.
- `-100 <= Node.val <= 100`

**Follow up:** Can you flatten the tree in-place (with `O(1)` extra space)?

## Solution

### Morris Traversal

对这个方法不怎么了解，需要再看。

Time: O(n), Space: O(n)

```shell
    1
   / \
  2   5
 / \   \
3   4   6

    1                   move left sub-tree of 1 to right
     \
      2     5
     / \     \
    3   4     6

    1                   insert right sub-tree of 1 to rightmost node
     \
      2
     / \
    3   4
         \
          5
           \
            6

    1                   move left sub-tree of 2 to right
     \
      2     4
       \     \
        3     5
               \
                6

    1                   insert right sub-tree of 2 to rightmost node
     \
      2
       \
        3
         \
          4
           \
            5
             \
              6
...
```

Refer:

- [Wikipedia/Morris-traversal-using-threading](https://en.wikipedia.org/wiki/Tree_traversal#Morris_in-order_traversal_using_threading)
- [leetcode/binary-tree-inorder-traversal/morris-traversal](https://leetcode.com/problems/binary-tree-inorder-traversal/solution/)
- [详细通俗的思路分析，多解法](https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list/solution/xiang-xi-tong-su-de-si-lu-fen-xi-duo-jie-fa-by--26/)
