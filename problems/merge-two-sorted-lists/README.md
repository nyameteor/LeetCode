# 21. Merge Two Sorted Lists

- Difficulty: Easy
- Topics: Linked List, Recursion
- Link: https://leetcode.com/problems/merge-two-sorted-lists/

## Description

Merge two sorted linked lists and return it as a **sorted** list. The list should be made by splicing together the nodes of the first two lists.

**Example 1:**

![img](https://assets.leetcode.com/uploads/2020/10/03/merge_ex1.jpg)

```
Input: l1 = [1,2,4], l2 = [1,3,4]
Output: [1,1,2,3,4,4]
```

**Example 2:**

```
Input: l1 = [], l2 = []
Output: []
```

**Example 3:**

```
Input: l1 = [], l2 = [0]
Output: [0]
```

**Constraints:**

- The number of nodes in both lists is in the range `[0, 50]`.
- `-100 <= Node.val <= 100`
- Both `l1` and `l2` are sorted in **non-decreasing** order.

## Solution

### Iteration

- We can use pointer `prev` to construct the result, and use pointer `l1` to iterate linked list `l1`, use `l2` to iterate linked list `l2`.
- To avoid the edge case, we can add a dummy pointer before the result pointer.
