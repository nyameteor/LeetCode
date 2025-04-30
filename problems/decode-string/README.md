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

### Approach: Stack-Based Decoding

- Use two stacks:
  - One for repeat counts (`countStack`)
  - One for accumulated strings (`strStack`)
- Iterate through the string:
  - If it's a digit, build the current repeat count.
  - If it's `'['`, push the current count and string onto the stacks, then reset them.
  - If it's `']'`, pop from both stacks and append the repeated substring.
  - Otherwise, append the character to the current string.

### References

- [[C++/Python] Clean & Simple Solutions w/ Explanation | Recursive + Iterative using Custom Stack](https://leetcode.com/problems/decode-string/solutions/1635464/c-python-clean-simple-solutions-w-explanation-recursive-iterative-using-custom-stack/)
