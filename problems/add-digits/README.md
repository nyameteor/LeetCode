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

### Number Theory: Digital Root

The **digital root** of a natural number is the single-digit value obtained by repeatedly summing its digits until only one digit remains.

In base $b$, the digital root $dr_b(n)$ can be computed directly using modular arithmetic:

$$
dr_b(n) =
    \begin{cases}
        0 & n = 0 \\
        1 + ((n-1)mod(b-1)) & n \neq 0
    \end{cases}
$$

### Simulation Approach

A straightforward implementation recursively sums the digits until the result is less than 10.

### References

- [Digital root](https://en.wikipedia.org/wiki/Digital_root)
