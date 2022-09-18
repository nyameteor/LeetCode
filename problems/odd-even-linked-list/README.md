# 328. Odd Even Linked List

- Difficulty: Medium
- Topics: Linked List
- Link: https://leetcode.com/problems/odd-even-linked-list/

## Description

Given the `head` of a singly linked list, group all the nodes with odd indices together followed by the nodes with even indices, and return _the reordered list_.

The **first** node is considered **odd**, and the **second** node is **even**, and so on.

Note that the relative order inside both the even and odd groups should remain as it was in the input.

You must solve the problem in `O(1)` extra space complexity and `O(n)` time complexity.

**Example 1:**

![img](https://assets.leetcode.com/uploads/2021/03/10/oddeven-linked-list.jpg)

```
Input: head = [1,2,3,4,5]
Output: [1,3,5,2,4]
```

**Example 2:**

![img](https://assets.leetcode.com/uploads/2021/03/10/oddeven2-linked-list.jpg)

```
Input: head = [2,1,3,5,6,4,7]
Output: [2,3,6,7,1,5,4]
```

**Constraints:**

- `n == `number of nodes in the linked list
- `0 <= n <= 10^4`
- `-10^6 <= Node.val <= 10^6`

## Solution

### In-place Transformation

直观的链表题，纸上画一下应该就能了解就地变换的过程。

可以参考 @archit91 的这张图([来源](<https://leetcode.com/problems/odd-even-linked-list/discuss/1607746/C%2B%2B-Simple-Solution-w-Images-and-Explanation-or-Brute-Force-%2B-O(1)-In-place-Transformation>))：

![](https://assets.leetcode.com/users/images/9005f9bb-e615-4737-b2cd-e578d595902e_1638439350.6554377.png)

感觉这个写法具有对称之美，比 `odd->next = odd->next->next` 的写法更优雅。

even 指针在每次循环结束时比 odd 走的更远，因此终止条件只需判断 even 和 even->next 是否为 Null。
