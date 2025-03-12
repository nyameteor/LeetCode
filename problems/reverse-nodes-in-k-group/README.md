# 25. Reverse Nodes in k-Group

- Difficulty: Hard
- Topics: Linked List, Recursion
- Link: https://leetcode.com/problems/reverse-nodes-in-k-group/

## Description

Given a linked list, reverse the nodes of a linked list _k_ at a time and return its modified list.

_k_ is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of _k_ then left-out nodes, in the end, should remain as it is.

You may not alter the values in the list's nodes, only nodes themselves may be changed.

**Example 1:**

![img](https://assets.leetcode.com/uploads/2020/10/03/reverse_ex1.jpg)

```
Input: head = [1,2,3,4,5], k = 2
Output: [2,1,4,3,5]
```

**Example 2:**

![img](https://assets.leetcode.com/uploads/2020/10/03/reverse_ex2.jpg)

```
Input: head = [1,2,3,4,5], k = 3
Output: [3,2,1,4,5]
```

**Example 3:**

```
Input: head = [1,2,3,4,5], k = 1
Output: [1,2,3,4,5]
```

**Example 4:**

```
Input: head = [1], k = 1
Output: [1]
```

**Constraints:**

- The number of nodes in the list is in the range `sz`.
- `1 <= sz <= 5000`
- `0 <= Node.val <= 1000`
- `1 <= k <= sz`

**Follow-up:** Can you solve the problem in O(1) extra memory space?

## Solution

### Recursive Approach

1. Count k nodes: Traverse the list to check if there are at least k nodes.
2. Reverse first k nodes: If we have k nodes, reverse them using a helper function.
3. Recursively process the remaining list: The next node after reversal points to the result of the recursive call.
4. Base case: If fewer than k nodes remain, return the head as is.

References: https://leetcode.com/problems/reverse-nodes-in-k-group/solutions/523641/c-iterative-solution-recursive-solution/
