# 258. Add Digits

- Difficulty: Easy
- Topics: Math, Simulation, Number Theory
- Link: https://leetcode.com/problems/add-digits/

## Description

Given an integer `num`, repeatedly add all its digits until the result has only one digit, and return it.

**Example 1:**

```
Input: num = 38
Output: 2
Explanation: The process is
38 --> 3 + 8 --> 11
11 --> 1 + 1 --> 2
Since 2 has only one digit, return it.
```

**Example 2:**

```
Input: num = 0
Output: 0
```

**Constraints:**

- `0 <= num <= 2^31 - 1`

**Follow up:** Could you do it without any loop/recursion in `O(1)` runtime?

## Solution

### Number Theory

Refer: https://en.wikipedia.org/wiki/Digital_root

The digital root (also repeated digital sum) of a natural number in a given radix is the (single digit) value obtained by an iterative process of summing digits, on each iteration using the result from the previous iteration to compute a digit sum. The process continues until a single-digit number is reached.

在数学中，数根是自然数的一种性质。数根是将一自然数的各个位数相加，若得到的值大于底数(base, radix)，则继续将各位数相加，直到其值小于底数为止。

Congruence formula

$$
dr_b(n) =
    \begin{cases}
        0 & n = 0 \\
        1 + ((n-1)mod(b-1)) & n \neq 0
    \end{cases}
$$

### Simulation

Naive implementation, pass.
