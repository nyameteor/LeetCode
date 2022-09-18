# 9. Palindrome Number

- Difficulty: Easy
- Topics: Math
- Link: https://leetcode.com/problems/palindrome-number/

## Description

Given an integer `x`, return `true` if `x` is palindrome integer.

An integer is a **palindrome** when it reads the same backward as forward. For example, `121` is palindrome while `123` is not.

**Example 1:**

```
Input: x = 121
Output: true
```

**Example 2:**

```
Input: x = -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
```

**Example 3:**

```
Input: x = 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
```

**Example 4:**

```
Input: x = -101
Output: false
```

**Constraints:**

- `-2^31 <= x <= 2^31 - 1`

**Follow up:** Could you solve it without converting the integer to a string?

## Solution

### Convert to String/Array

将 int 转换为 string/array 后，使用头尾两个指针同时向中心遍历，通过头尾索引对应元素是否相等判断回文，普通直观的方法。

### Revert half of the number

Todo
