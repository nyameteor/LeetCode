# 50. Pow(x, n)

- Difficulty: Medium
- Topics: Math, Recursion
- Link: https://leetcode.com/problems/powx-n/

## Description

Implement [pow(x, n)](http://www.cplusplus.com/reference/valarray/pow/), which calculates `x` raised to the power `n` (i.e., `xn`).

**Example 1:**

```
Input: x = 2.00000, n = 10
Output: 1024.00000
```

**Example 2:**

```
Input: x = 2.10000, n = 3
Output: 9.26100
```

**Example 3:**

```
Input: x = 2.00000, n = -2
Output: 0.25000
Explanation: 2^-2 = 1/2^2 = 1/4 = 0.25
```

**Constraints:**

- `-100.0 < x < 100.0`
- `-2^31 <= n <= 2^31-1`
- `-10^4 <= x^n <= 10^4`

## Solution

We can use **exponentiation by squaring** to efficiently compute `x^n` for both positive and negative `n`. The approach recursively reduces the problem size by halving the exponent in each step:

- For even `n`, the result is computed as `half * half`, where `half = x^(n/2)`.
- For odd `n`, it is `x * half * half`, which includes an extra multiplication by `x` for the odd exponent.

For negative exponents, we compute the reciprocal (`1 / x^(-n)`), ensuring correct handling of large values using `int64` for the exponent.

Time Complexity: O(log n)

References: https://en.wikipedia.org/wiki/Exponentiation_by_squaring
