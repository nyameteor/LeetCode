# 1015. Smallest Integer Divisible by K

- Difficulty: Medium
- Topics: Hash Table, Math
- Link: https://leetcode.com/problems/smallest-integer-divisible-by-k/

## Description

Given a positive integer `k`, you need to find the **length** of the **smallest** positive integer `n` such that `n` is divisible by `k`, and `n` only contains the digit `1`.

Return _the **length** of_ `n`. If there is no such `n`, return -1.

**Note:** `n` may not fit in a 64-bit signed integer.

**Example 1:**

```
Input: k = 1
Output: 1
Explanation: The smallest answer is n = 1, which has length 1.
```

**Example 2:**

```
Input: k = 2
Output: -1
Explanation: There is no such positive integer n divisible by 2.
```

**Example 3:**

```
Input: k = 3
Output: 3
Explanation: The smallest answer is n = 111, which has length 3.
```

**Constraints:**

- `1 <= k <= 10^5`

## Solution

### Key Idea

- We're looking for the smallest number made only of `1`s (like `1`, `11`, `111`, ...) that is divisible by `k`.
- Instead of building the number, we can simulate the remainder:
  `newRemainder = (remainder * 10 + 1) % k`
- Stop when remainder becomes `0` (found) or repeats (no solution).

### Optimization

- If `k % 2 == 0` or `k % 5 == 0`, return `-1` early (repunit numbers are never divisible by 2 or 5).

### References

- [[C++/Python] Simple Solution w/ Explanation | Brute+ Optimizations w/ Math & Pigeonhole](https://leetcode.com/problems/smallest-integer-divisible-by-k/solutions/1656121/c-python-simple-solution-w-explanation-brute-optimizations-w-math-pigeonhole/)
- [Pigeonhole principle](https://en.wikipedia.org/wiki/Pigeonhole_principle)
