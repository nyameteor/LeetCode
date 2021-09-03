# 415. Add Strings

- Difficulty: Easy
- Topics: Math, String, Simulation
- Link: https://leetcode.com/problems/add-strings/

## Description

Given two non-negative integers, `num1` and `num2` represented as string, return _the sum of_ `num1` _and_ `num2` _as a string_.

You must solve the problem without using any built-in library for handling large integers (such as `BigInteger`). You must also not convert the inputs to integers directly.

**Example 1:**

```
Input: num1 = "11", num2 = "123"
Output: "134"
```

**Example 2:**

```
Input: num1 = "456", num2 = "77"
Output: "533"
```

**Example 3:**

```
Input: num1 = "0", num2 = "0"
Output: "0"
```

**Constraints:**

- `1 <= num1.length, num2.length <= 10^4`
- `num1` and `num2` consist of only digits.
- `num1` and `num2` don't have any leading zeros except for the zero itself.

## Solution

给定两个非负整数，`num1` 和 `num2` 表示为字符串，返回 _`num1`_ 和 _`num2`_ 的和（表示为字符串）。

简单的模拟题，不能直接用 String 转 Integer 方法，所以需要模拟加法的过程，从右到左对数字逐个相加，并处理进位。

注意 char 和 int 的互相转换，Java 中 int 和 char 互相转换示例：

```java
int a = 1;
char b = '2';

// int to char
char c = (char) (a + '0');

// char to int
int d = b - '0';
```

注意在循环体内，字符串的联接应当通过 StringBuilder 进行。

原因：反编译出的字节码文件显示使用 `+` 的字符串联接每次循环都会 new 一个 StringBuilder 对象然后操作，最后通过 toString 方法返回 String 对象，造成内存资源浪费并降低运行速度。
