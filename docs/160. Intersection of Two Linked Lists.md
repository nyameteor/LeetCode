# 160. Intersection of Two Linked Lists

- Difficulty: Easy
- Topics: Hash Table, Linked List, Two Pointers
- Link: https://leetcode.com/problems/intersection-of-two-linked-lists/

## Description

Given the heads of two singly linked-lists `headA` and `headB`, return _the node at which the two lists intersect_. If the two linked lists have no intersection at all, return `null`.

For example, the following two linked lists begin to intersect at node `c1`:

![img](https://assets.leetcode.com/uploads/2021/03/05/160_statement.png)

The test cases are generated such that there are no cycles anywhere in the entire linked structure.

**Note** that the linked lists must **retain their original structure** after the function returns.

**Custom Judge:**

The inputs to the **judge** are given as follows (your program is **not** given these inputs):

- `intersectVal` - The value of the node where the intersection occurs. This is `0` if there is no intersected node.
- `listA` - The first linked list.
- `listB` - The second linked list.
- `skipA` - The number of nodes to skip ahead in `listA` (starting from the head) to get to the intersected node.
- `skipB` - The number of nodes to skip ahead in `listB` (starting from the head) to get to the intersected node.

The judge will then create the linked structure based on these inputs and pass the two heads, `headA` and `headB` to your program. If you correctly return the intersected node, then your solution will be **accepted**.

**Example 1:**

![img](https://assets.leetcode.com/uploads/2021/03/05/160_example_1_1.png)

```
Input: intersectVal = 8, listA = [4,1,8,4,5], listB = [5,6,1,8,4,5], skipA = 2, skipB = 3
Output: Intersected at '8'
Explanation: The intersected node's value is 8 (note that this must not be 0 if the two lists intersect).
From the head of A, it reads as [4,1,8,4,5]. From the head of B, it reads as [5,6,1,8,4,5]. There are 2 nodes before the intersected node in A; There are 3 nodes before the intersected node in B.
```

**Example 2:**

![img](https://assets.leetcode.com/uploads/2021/03/05/160_example_2.png)

```
Input: intersectVal = 2, listA = [1,9,1,2,4], listB = [3,2,4], skipA = 3, skipB = 1
Output: Intersected at '2'
Explanation: The intersected node's value is 2 (note that this must not be 0 if the two lists intersect).
From the head of A, it reads as [1,9,1,2,4]. From the head of B, it reads as [3,2,4]. There are 3 nodes before the intersected node in A; There are 1 node before the intersected node in B.
```

**Example 3:**

![img](https://assets.leetcode.com/uploads/2021/03/05/160_example_3.png)

```
Input: intersectVal = 0, listA = [2,6,4], listB = [1,5], skipA = 3, skipB = 2
Output: No intersection
Explanation: From the head of A, it reads as [2,6,4]. From the head of B, it reads as [1,5]. Since the two lists do not intersect, intersectVal must be 0, while skipA and skipB can be arbitrary values.
Explanation: The two lists do not intersect, so return null.
```

**Constraints:**

- The number of nodes of `listA` is in the `m`.
- The number of nodes of `listB` is in the `n`.
- `0 <= m, n <= 3 * 10^4`
- `1 <= Node.val <= 10^5`
- `0 <= skipA <= m`
- `0 <= skipB <= n`
- `intersectVal` is `0` if `listA` and `listB` do not intersect.
- `intersectVal == listA[skipA] == listB[skipB]` if `listA` and `listB` intersect.

**Follow up:** Could you write a solution that runs in `O(n)` time and use only `O(1)` memory?

## Solution

编写一个程序，找到两个单链表相交的起始节点，若两个链表没有相交则返回 null。

## Two Pointers, One-Pass

非常巧妙的方法，使用两个指针遍历这两个链表，当指针遍历到列表末尾时，从另一个链表的头开始遍历。这确保了两个指针会同时到达交点（或末尾）。

设 listA 的长度为 `m`，listB 的长度为 `n`。

指针 `pA` 和 `pB` 从各自链表的头开始遍历，直到 `pA == pB` 时退出：

- 若两个链表存在交点，在两个指针走完 `m+n` 个位置前，两个指针一定会在某个位置相等，退出循环；
- 若两个链表不存在交点，在两个指针走完 `m+n` 个位置后，两个指针都为 null（同时到达链表末尾），退出循环；
- 返回该指针。

Example 2:

```shell
0 -> 9 -> 1 -> 2 -> 4 -> 3 -> 2 -> 4
                    * tailA
                              * 交点
          * tailB
3 -> 2 -> 4 -> 0 -> 9 -> 1 -> 2 -> 4
```

Example 3:

```shell
2 -> 6 -> 4 -> 1 -> 5 -> null
          * tailA
                         * 无交点
     * tailB
1 -> 5 -> 2 -> 6 -> 4 -> null
```

**References**

- https://leetcode.com/problems/intersection-of-two-linked-lists/discuss/2116221/Visual-Explanation-or-One-Pass-JAVA
