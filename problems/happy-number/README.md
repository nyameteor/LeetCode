# 202. Happy Number

- Difficulty: Easy
- Topics: Hash Table, Math, Two Pointers
- Link: https://leetcode.com/problems/happy-number/

## Description

Write an algorithm to determine if a number `n` is happy.

A **happy number** is a number defined by the following process:

- Starting with any positive integer, replace the number by the sum of the squares of its digits.
- Repeat the process until the number equals 1 (where it will stay), or it **loops endlessly in a cycle** which does not include 1.
- Those numbers for which this process **ends in 1** are happy.

Return `true` _if_ `n` _is a happy number, and_ `false` _if not_.

**Example 1:**

```
Input: n = 19
Output: true
Explanation:
12 + 92 = 82
82 + 22 = 68
62 + 82 = 100
12 + 02 + 02 = 1
```

**Example 2:**

```
Input: n = 2
Output: false
```

**Constraints:**

- `1 <= n <= 2^31 - 1`

## Solution

We can use a HashSet to detect cycles, or [Floyd's Cycle Detection](https://en.wikipedia.org/wiki/Cycle_detection#Floyd's_tortoise_and_hare) (slow/fast pointers) for constant space.
