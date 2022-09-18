# 86. Partition List

- Difficulty: Medium
- Topics: Linked List, Two Pointers
- Link: https://leetcode.com/problems/partition-list/

## Description

Given the `head` of a linked list and a value `x`, partition it such that all nodes **less than** `x` come before nodes **greater than or equal** to `x`.

You should **preserve** the original relative order of the nodes in each of the two partitions.

**Example 1:**

![](https://assets.leetcode.com/uploads/2021/01/04/partition.jpg)

```
Input: head = [1,4,3,2,5,2], x = 3
Output: [1,2,2,4,3,5]
```

**Example 2:**

```
Input: head = [2,1], x = 2
Output: [1,2]
```

**Constraints:**

- The number of nodes in the list is in the range `[0, 200]`.
- `-100 <= Node.val <= 100`
- `-200 <= x <= 200`

## Solution

We can create two smaller linked lists during traversing original linked list:

- `Before`: all the elements smaller than `x`
- `After`: all the elements equal to or greater than `x`.

Simply join `Before` and `After`, then we get the reformed list.

**References:**

- https://leetcode.com/problems/partition-list/solution/
