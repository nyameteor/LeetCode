# 109. Convert Sorted List to Binary Search Tree

- Difficulty: Medium
- Topics: Linked List, Divide and Conquer, Tree, Binary Search Tree, Binary Tree
- Link: https://leetcode.com/problems/convert-sorted-list-to-binary-search-tree/

## Description

Given the `head` of a singly linked list where elements are sorted in **ascending order**, convert *it to a* ***height-balanced*** *binary search tree*.

**Example 1:**

![](https://assets.leetcode.com/uploads/2020/08/17/linked.jpg)

```
**Input:** head = [-10,-3,0,5,9]
**Output:** [0,-3,9,-10,null,5]
**Explanation:** One possible answer is [0,-3,9,-10,null,5], which represents the shown height balanced BST.

```

**Example 2:**

```
**Input:** head = []
**Output:** []

```

**Constraints:**

- The number of nodes in `head` is in the range `[0, 2 * 10^4]`.
- `-10^5 <= Node.val <= 10^5`

## Solution

We can set the middle element of the sorted list as the root, and perform this operation recursively for the left half and the right half to build the BST.

To do get the middle element of the linked list, we can convert it to an array, or we can use fast and slow pointers.
