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

### Brute Force

暴力计算的时间复杂度为 O(n) 会超时。

关于快速幂：[Quick Pow](https://oi-wiki.org/math/quick-pow/)

### Recursion

使用 Recursion，递推方程如下：

```
// i % 2 == 0
Pow(x, i) = Pow(x * x, i/2)

// i % 2 == 1
Pow(x, i) = Pow(x * x, (i-1)/2) * x
```

n < 0 时，求 `Pow(x,n)` 可以转换为求 `1 / Pow(x, -n)`。

另外，`n = -2^31(INT_MIN)` 为 edge case，此时直接将 `n` 赋值为 `-n` 会溢出，需要单独处理。

- Time: O(log(n))
