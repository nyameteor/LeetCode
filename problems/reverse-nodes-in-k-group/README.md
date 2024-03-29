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

给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。

翻转链表并不麻烦，可以参考 [206. Reverse Linked List](./docs/206.%20Reverse%20Linked%20List.md)。但本题需要按照 k 进行分组，需要考虑的细节比较多：

- 按照 k 分组，每次翻转前判断剩余的节点是否大于 k，若小于 k 则不翻转
- 翻转子链表时，翻转自身后还需要将子链表的 head 与上一个子链表连接，将 tail 与下一个子链表连接。第一个子链表的 head 前没有节点，为了防止特判可以加一个 dummy head

```shell
k = 3

------------------- init ----------------------
# add dummy head to avoid corner case
# d: dummy head

 d
 |
-1 -> 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7

----------- reverse first group ---------------
# p: sub list previous, n: sub list next

-1 -> 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7       p = d, t = d
p,t

-1 -> 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7       t = t.next * k
p               t

        sub list to be reversed
           |
      h ------- t
      |         |
-1 -> 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7       h = p.next, n = t.next
 p                   n

# revese sub list, return new head and tail
# h: head, t: tail, new head will be `t`, new tail be `h`
# b: backfoward, c: current, f: foward

1 -> 2 -> 3
h         t

1 -> 2 -> 3         b = head, c = b.next, f = b.next
b   c,f
----------------
1 <- 2    3         f = c.next, c.next = b
b    c    f

1 <- 2    3         b = c, c = f
     b   c,f
-----------------
1 <- 2 <- 3         f = c.next, c.next = b
     b    c    f

1 <- 2 <- 3         b = c, c = f
          b   c,f
-----------------   b == t, return new head, new tail

# connect sub list head and tail
-1    3 -> 2 -> 1    4 -> 5 -> 6 -> 7
 p    h         t    n

-1 -> 3 -> 2 -> 1 -> 4 -> 5 -> 6 -> 7
 p    h         t    n

-1 -> 3 -> 2 -> 1 -> 4 -> 5 -> 6 -> 7       p = t
               p,t

---------- reverse second group ---------------

-1 -> 3 -> 2 -> 1 -> 4 -> 5 -> 6 -> 7
               p,t
...
```

更多内容可以参考[官方题解](https://leetcode-cn.com/problems/reverse-nodes-in-k-group/solution/k-ge-yi-zu-fan-zhuan-lian-biao-by-leetcode-solutio/)。
