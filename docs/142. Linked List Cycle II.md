# 142. Linked List Cycle II

- Difficulty: Medium
- Topics: Hash Table, Linked List, Two Pointers
- Link: https://leetcode.com/problems/linked-list-cycle-ii/

## Description

Given the `head` of a linked list, return _the node where the cycle begins. If there is no cycle, return_ `null`.

There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the `next` pointer. Internally, `pos` is used to denote the index of the node that tail's `next` pointer is connected to (**0-indexed**). It is `-1` if there is no cycle. **Note that** `pos` **is not passed as a parameter**.

**Do not modify** the linked list.

**Example 1:**

![img](https://assets.leetcode.com/uploads/2018/12/07/circularlinkedlist.png)

```
Input: head = [3,2,0,-4], pos = 1
Output: tail connects to node index 1
Explanation: There is a cycle in the linked list, where tail connects to the second node.
```

**Example 2:**

![img](https://assets.leetcode.com/uploads/2018/12/07/circularlinkedlist_test2.png)

```
Input: head = [1,2], pos = 0
Output: tail connects to node index 0
Explanation: There is a cycle in the linked list, where tail connects to the first node.
```

**Example 3:**

![img](https://assets.leetcode.com/uploads/2018/12/07/circularlinkedlist_test3.png)

```
Input: head = [1], pos = -1
Output: no cycle
Explanation: There is no cycle in the linked list.
```

**Constraints:**

- The number of the nodes in the list is in the range `[0, 10^4]`.
- `-10^5 <= Node.val <= 10^5`
- `pos` is `-1` or a **valid index** in the linked-list.

**Follow up:** Can you solve it using `O(1)` (i.e. constant) memory?

## Solution

### Two Pointers, Brute Force (Accepted)

暴力解法，对链表中的每个节点依次使用“双指针检查环”算法，直到找到某个节点，快慢指针相遇的位置为该节点，说明该节点是环的开始节点。

Time Complexity: O(n^2)

### Two Pointers, Two Pass

对链表使用“双指针检查环”算法，而后将慢指针移回 head，之后两个原本的快慢指针以相同的速度移动（每次一步），双指针再次相遇的地方即是环的开始节点。

该解法只需两躺遍历。

Time Complexity: O(n)
