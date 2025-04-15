# 234. Palindrome Linked List

- Difficulty: Easy
- Topics: Linked List, Two Pointers, Stack, Recursion
- Link: https://leetcode.com/problems/palindrome-linked-list/

## Description

Given the `head` of a singly linked list, return `true` if it is a palindrome.

**Example 1:**

![img](https://assets.leetcode.com/uploads/2021/03/03/pal1linked-list.jpg)

```
Input: head = [1,2,2,1]
Output: true
```

**Example 2:**

![img](https://assets.leetcode.com/uploads/2021/03/03/pal2linked-list.jpg)

```
Input: head = [1,2]
Output: false
```

**Constraints:**

- The number of nodes in the list is in the range `[1, 10^5]`.
- `0 <= Node.val <= 9`

**Follow up:** Could you do it in `O(n)` time and `O(1)` space?

## Solution

### Approach: Reverse Second Half

- Find the middle of the linked list using slow and fast pointers.
- Reverse the second half of the linked list in-place.
- Compare the first half with the reversed second half node by node.
