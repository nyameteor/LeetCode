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

Power of 2 means only one bit of `n` is 1.

### Count bits

Count the number of bits 1, if it only has one, then `n` is a power of two.

### Math & Bit Hack

We can use `n & (n - 1)` to figure out if `n` is either 0 or an exact power of two.

```shell
  1000 0000 0000 0000
&  111 1111 1111 1111
  ==== ==== ==== ====
= 0000 0000 0000 0000
```

Refer:

- https://stackoverflow.com/a/4678376
- https://github.com/haoel/leetcode/blob/master/algorithms/cpp/powerOfTwo/PowerOfTwo.cpp

More Bit hacks: http://graphics.stanford.edu/~seander/bithacks.html
