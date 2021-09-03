# 206. Reverse Linked List

- Difficulty: Easy
- Topics: Linked List, Recursion
- Link: https://leetcode.com/problems/reverse-linked-list/

## Description

Given the `head` of a singly linked list, reverse the list, and return _the reversed list_.

**Example 1:**

![img](https://assets.leetcode.com/uploads/2021/02/19/rev1ex1.jpg)

```
Input: head = [1,2,3,4,5]
Output: [5,4,3,2,1]
```

**Example 2:**

![img](https://assets.leetcode.com/uploads/2021/02/19/rev1ex2.jpg)

```
Input: head = [1,2]
Output: [2,1]
```

**Example 3:**

```
Input: head = []
Output: []
```

**Constraints:**

- The number of nodes in the list is the range `[0, 5000]`.
- `-5000 <= Node.val <= 5000`

**Follow up:** A linked list can be reversed either iteratively or recursively. Could you implement both?

## Solution

### Iteration

遍历迭代，通过三个遍历指针 `backword`, `current`, `forward` 将链表元素反转（通过改变 `current.next` 的值）。

Example 1:

```shell
------ init --------

1 -> 2 -> 3 -> 4 -> 5       b = h, c = h.next, f = c
b   c,f

------ loop --------

-------- 1 ---------

1 -> 2 -> 3 -> 4 -> 5       f = c.next
b    c    f

1 <- 2    3 -> 4 -> 5       c.next = b
b    c    f

1 <- 2    3 -> 4 -> 5       b = c, c = f
     b   f,c

-------- 2 ---------

1 <- 2    3 -> 4 -> 5       f = c.next
     b    c    f

1 <- 2 <- 3    4 -> 5       c.next = b
     b    c    f

1 <- 2 <- 3    4 -> 5       b = c, c = f
          b   f,c

...
```
