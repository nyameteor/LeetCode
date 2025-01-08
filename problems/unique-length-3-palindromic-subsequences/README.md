# 1930. Unique Length-3 Palindromic Subsequences

- Difficulty: Medium
- Topics: Hash Table, String, Bit Manipulation, Prefix Sum
- Link: https://leetcode.com/problems/unique-length-3-palindromic-subsequences/

## Description

Given a string `s`, return *the number of **unique palindromes of length three** that are a **subsequence** of* `s`.

Note that even if there are multiple ways to obtain the same subsequence, it is still only counted **once**.

A **palindrome** is a string that reads the same forwards and backwards.

A **subsequence** of a string is a new string generated
from the original string with some characters (can be none) deleted
without changing the relative order of the remaining characters.

- For example, `"ace"` is a subsequence of `"abcde"`.

**Example 1:**

```
**Input:** s = "aabca"
**Output:** 3
**Explanation:** The 3 palindromic subsequences of length 3 are:
- "aba" (subsequence of "aabca")
- "aaa" (subsequence of "aabca")
- "aca" (subsequence of "aabca")

```

**Example 2:**

```
**Input:** s = "adc"
**Output:** 0
**Explanation:** There are no palindromic subsequences of length 3 in "adc".

```

**Example 3:**

```
**Input:** s = "bbcbaba"
**Output:** 4
**Explanation:** The 4 palindromic subsequences of length 3 are:
- "bbb" (subsequence of "bbcbaba")
- "bcb" (subsequence of "bbcbaba")
- "bab" (subsequence of "bbcbaba")
- "aba" (subsequence of "bbcbaba")

```

**Constraints:**

- `3 <= s.length <= 10^5`
- `s` consists of only lowercase English letters.

## Solution

### Approach: Character Tracking

1. A palindrome of length 3 has the form `x_y_x`, where `x` is the first and last character, and `y` is the middle character.
2. To count unique palindromic subsequences, we can:
    - For each unique character `x` (the first and last character of the palindrome), determine the range `[l, r]` where `x` occurs in the string.
    - Track all unique characters `y` that appear between `l` and `r`.
3. Use a set to store unique palindromes formed by `x_y_x`.
