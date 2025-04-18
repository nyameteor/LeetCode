# 328. Odd Even Linked List

- Difficulty: Medium
- Topics: Linked List
- Link: https://leetcode.com/problems/odd-even-linked-list/

## Description

Given the `head` of a singly linked list, group all the nodes with odd indices together followed by the nodes with even indices, and return *the reordered list*.

The **first** node is considered **odd**, and the **second** node is **even**, and so on.

Note that the relative order inside both the even and odd groups should remain as it was in the input.

You must solve the problem in `O(1)` extra space complexity and `O(n)` time complexity.

**Example 1:**

![](https://assets.leetcode.com/uploads/2021/03/10/oddeven-linked-list.jpg)

```
Input: head = [1,2,3,4,5]
Output: [1,3,5,2,4]
```

**Example 2:**

![](https://assets.leetcode.com/uploads/2021/03/10/oddeven2-linked-list.jpg)

```
Input: head = [2,1,3,5,6,4,7]
Output: [2,3,6,7,1,5,4]
```

**Constraints:**

- The number of nodes in the linked list is in the range `[0, 10^4]`.
- `-10^6 <= Node.val <= 10^6`

## Solution

### Approach: Two Pointers

Group nodes by index parity in a single pass using two pointers:

- Use `odd` and `even` pointers to build two separate chains.
- Link all odd-indexed nodes together, and all even-indexed nodes together.
- Finally, connect the tail of the odd chain to the head of the even chain.

This preserves the relative order and runs in **O(n)** time with **O(1)** space.

### References

- [[C++] Simple Solution w/ Images & Explanation](https://leetcode.com/problems/odd-even-linked-list/solutions/1607746/c-simple-solution-w-images-explanation-brute-force-o-1-in-place-transformation/)
