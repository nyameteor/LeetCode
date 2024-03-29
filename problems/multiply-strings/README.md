# 43. Multiply Strings

- Difficulty: Medium
- Topics: Math, String, Simulation
- Link: https://leetcode.com/problems/multiply-strings/

## Description

Given two non-negative integers `num1` and `num2` represented as strings, return the product of `num1` and `num2`, also represented as a string.

**Note:** You must not use any built-in BigInteger library or convert the inputs to integer directly.

**Example 1:**

```
Input: num1 = "2", num2 = "3"
Output: "6"
```

**Example 2:**

```
Input: num1 = "123", num2 = "456"
Output: "56088"
```

**Constraints:**

- `1 <= num1.length, num2.length <= 200`
- `num1` and `num2` consist of digits only.
- Both `num1` and `num2` do not contain any leading zero, except the number `0` itself.

## Solution

本题正要求实现大数（使用字符串表示）的相乘，要求不使用内建的 Biginter 库或者将字符串转换为数字（会溢出）再相乘。

一些语言（如 Python）中，也会将超出类型上界的大数转换为字符串，以便进行后续运算。

官方题解：https://leetcode.com/problems/multiply-strings/solution/

### Elementary Math

简单地模拟乘法计算的过程。可以用一些 C++ 字符处理的技巧（见代码）。

![slide1](https://leetcode.com/problems/multiply-strings/Figures/43/Slide1.JPG)

![slide2](https://leetcode.com/problems/multiply-strings/Figures/43/Slide2.JPG)

![slide3](https://leetcode.com/problems/multiply-strings/Figures/43/Slide3.JPG)
