# 394. Decode String

- Difficulty: Medium
- Topics: String, Stack, Recursion
- Link: https://leetcode.com/problems/decode-string/

## Description

Given an encoded string, return its decoded string.

The encoding rule is: `k[encoded_string]`, where the `encoded_string` inside the square brackets is being repeated exactly `k` times. Note that `k` is guaranteed to be a positive integer.

You may assume that the input string is always valid; No extra white spaces, square brackets are well-formed, etc.

Furthermore, you may assume that the original data does not contain any digits and that digits are only for those repeat numbers, `k`. For example, there won't be input like `3a` or `2[4]`.

**Example 1:**

```
Input: s = "3[a]2[bc]"
Output: "aaabcbc"
```

**Example 2:**

```
Input: s = "3[a2[c]]"
Output: "accaccacc"
```

**Example 3:**

```
Input: s = "2[abc]3[cd]ef"
Output: "abcabccdcdcdef"
```

**Example 4:**

```
Input: s = "abc3[cd]xyz"
Output: "abccdcdcdxyz"
```

**Constraints:**

- `1 <= s.length <= 30`
- `s` consists of lowercase English letters, digits, and square brackets `'[]'`.
- `s` is guaranteed to be **a valid** input.
- All the integers in `s` are in the range `[1, 300]`.

## Solution

### Two Stack

和 **Basic Calculator** 相似的解法。使用两个栈，栈 a 存放操作的数字，栈 b 存放字符。

遍历字符串 s：

- 若 s[i] 是数字，则读取所有数字的字符，将其转换为 int 类型后推入栈 a。
- 若 s[i] 是字符 ']'：
  - 推出栈 b 中 '[' 和 ']' 间的字符，即为经过编码后的字符串。
  - 推出栈 a 的栈顶，即为解码的参数 k。
  - 对编码后的字符串进行解码（重复 k 次），将解码后的字符推入栈 b。
- 若 s[i] 是其他字符，则推入栈 b。

遍历结束后，推出栈 b 中的所有字符，并逆序相加，即为结果。

Tips:

- c++ 中判断 char `c` 是否为数字，可以使用 `isdigit(c)`，等价于 `c >= '0' && c <= '9'`。
