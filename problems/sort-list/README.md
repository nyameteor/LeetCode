# 148. Sort List

- Difficulty: Medium
- Topics: Linked List, Two Pointers, Divide and Conquer, Sorting, Merge Sort
- Link: https://leetcode.com/problems/sort-list/

## Description

Given the `head` of a linked list, return \*the list after sorting it in **ascending order\***.

**Example 1:**

![img](https://assets.leetcode.com/uploads/2020/09/14/sort_list_1.jpg)

```
Input: head = [4,2,1,3]
Output: [1,2,3,4]
```

**Example 2:**

![img](https://assets.leetcode.com/uploads/2020/09/14/sort_list_2.jpg)

```
Input: head = [-1,5,3,4,0]
Output: [-1,0,3,4,5]
```

**Example 3:**

```
Input: head = []
Output: []
```

**Constraints:**

- The number of nodes in the list is in the range `[0, 5 * 10^4]`.
- `-10^5 <= Node.val <= 10^5`

**Follow up:** Can you sort the linked list in `O(n logn)` time and `O(1)` memory (i.e. constant space)?

## Solution

To implement bottom-up merge sort with `O(n logn)` time and `O(1)` space, please refer to: [Bottom-up implementation](https://en.wikipedia.org/wiki/Merge_sort#Bottom-up_implementation)
