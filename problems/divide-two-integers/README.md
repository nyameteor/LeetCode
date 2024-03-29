# 29. Divide Two Integers

- Difficulty: Medium
- Topics: Math, Bit Manipulation
- Link: https://leetcode.com/problems/divide-two-integers/

## Description

Given two integers `dividend` and `divisor`, divide two integers without using multiplication, division, and mod operator.

Return the quotient after dividing `dividend` by `divisor`.

The integer division should truncate toward zero, which means losing its fractional part. For example, `truncate(8.345) = 8` and `truncate(-2.7335) = -2`.

**Note:** Assume we are dealing with an environment that could only store integers within the **32-bit** signed integer range: `[−2^31, 2^31 − 1]`. For this problem, assume that your function **returns** `2^31 − 1` **when the division result overflows**.

**Example 1:**

```
Input: dividend = 10, divisor = 3
Output: 3
Explanation: 10/3 = truncate(3.33333..) = 3.
```

**Example 2:**

```
Input: dividend = 7, divisor = -3
Output: -2
Explanation: 7/-3 = truncate(-2.33333..) = -2.
```

**Example 3:**

```
Input: dividend = 0, divisor = 1
Output: 0
```

**Example 4:**

```
Input: dividend = 1, divisor = 1
Output: 1
```

**Constraints:**

- `-2^31 <= dividend, divisor <= 2^31 - 1`
- `divisor != 0`

## Solution

给定两个为整数的被除数和除数，在不使用乘法、除法和 mod 运算符的情况下将两个整数相除，返回商。

## Recursion

使用快速幂的思路，本题中将加法改为减法：当除数翻倍，被除数减不下去的时候，递归尝试以较小的倍数继续减。

Base Case:

```shell
dividend - divisor  < 0, return 0
dividend - divisor == 0, return 1
```

递归方程式：

- （不知如何清楚地表达，待补充）

- Time: O(log(n))

这个问题比较坑的地方是除数、被除数、商都可能出现等于 INT_MIN 的边缘情况，如果每个都考虑会变得很复杂。

比较好的方式是将被除数与除数转换为 long 类型，再计算商；返回商的时候判断边缘情况（这样只需判断一次）。

### Bit Manipulation

Todo
