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

- The number of nodes in the list is in the range `[0, 5 * 104]`.
- `-105 <= Node.val <= 105`

**Follow up:** Can you sort the linked list in `O(n logn)` time and `O(1)` memory (i.e. constant space)?

## Solution

### Natural merge Sort

Refer: https://en.wikipedia.org/wiki/Merge_sort#Natural_merge_sort

Use natural merge sort to linked list. Split linked list into monotonic sub linked list, then merge them until there is only one sub linked list, which means the linked list is sorted.

Example:

```shell
Start       :  3  4  2  1  7  5  8  9  0  6
Select runs : (3  4)(2)(1  7)(5  8  9)(0  6)
Merge       : (2  3  4)(1  5  7  8  9)(0  6)
Merge       : (1  2  3  4  5  7  8  9)(0  6)
Merge       : (0  1  2  3  4  5  6  7  8  9)
```
