# 231. Power of Two

- Difficulty: Easy
- Topics: Math, Bit Manipulation, Recursion
- Link: https://leetcode.com/problems/power-of-two/

## Description

Given an integer `n`, return _`true` if it is a power of two. Otherwise, return `false`_.

An integer `n` is a power of two, if there exists an integer `x` such that `n == 2x`.

**Example 1:**

```
Input: n = 1
Output: true
Explanation: 20 = 1
```

**Example 2:**

```
Input: n = 16
Output: true
Explanation: 24 = 16
```

**Example 3:**

```
Input: n = 3
Output: false
```

**Constraints:**

- `-2^31 <= n <= 2^31 - 1`

## Solution

A number is a power of two if it has only one bit set.

### Bit Count Check

Count the number of 1 bits in n. If there's only one, it's a power of two.

### Bitwise Check

Use `n > 0 && (n & (n - 1)) == 0` to check:

```
  1000 0000
& 0111 1111
= 0000 0000
```

### References

- https://stackoverflow.com/a/4678376
- https://graphics.stanford.edu/~seander/bithacks.html
