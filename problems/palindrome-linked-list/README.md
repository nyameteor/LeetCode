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

### Use Array to simulate

简单直观的方法，把链表拷贝到一个数组中进行模拟。然后设置头尾指针向中间遍历来判断是否为回文。

Time: `O(n)`, Space: `O(n)`

### Reverse linked list

- 使用快慢指针（快指针一次走 2 steps，慢指针一次走 1 step）找到链表 h1 的中心节点。
- 逆转中心节点之后的节点（将后半段逆转）得到 h2。
- 使用 h1, h2 遍历，比较节点数值，若至比较结束节点数值都相等，则链表为回文。

```shell

find Middle
1 -> 2 -> 3 -> 2 -> 1 -> null       fast->next = null, middle = slow
         slow      fast

reverse second half of linked list from `mid`
1 -> 2 -> 3 -> 2 -> 1 -> null
         mid

1 -> 2 -> 3 <- 2 <- 1
p        mid        q

traverse list with `p`, `q` two pointers
1 -> 2 -> 3 <- 2 <- 1
p                   q           p->val = q->val
     p         q                p->val = q->val
         p,q                    p = mid, loop end
```

Ref: [haoel/leetcode/PalindromeLinkedList.cpp](https://github.com/haoel/leetcode/blob/master/algorithms/cpp/palindromeLinkedList/PalindromeLinkedList.cpp)

Time: `O(n)`, Space: `O(1)`
