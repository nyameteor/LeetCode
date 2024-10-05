# 567. Permutation in String

- Difficulty: Medium
- Topics: Hash Table, Two Pointers, String, Sliding Window
- Link: https://leetcode.com/problems/permutation-in-string/

## Description

Given two strings `s1` and `s2`, return `true` if `s2` contains a permutation of `s1`, or `false` otherwise.

In other words, return `true` if one of `s1`'s permutations is the substring of `s2`.

**Example 1:**

```
**Input:** s1 = "ab", s2 = "eidbaooo"
**Output:** true
**Explanation:** s2 contains one permutation of s1 ("ba").
```

**Example 2:**

```
**Input:** s1 = "ab", s2 = "eidboaoo"
**Output:** false
```

**Constraints:**

- `1 <= s1.length, s2.length <= 10^4`
- `s1` and `s2` consist of lowercase English letters.

## Solution

We can use a hashmap `s1Count` to store the character frequencies in `s1`. To check for permutations of `s1` within `s2`, we examine each possible substring of `s2` that has the same length as `s1` by maintaining a corresponding hashmap `s2Count`.

Using a sliding window of size `len(s1)`, we iterate through `s2`. Whenever `s1Count` matches `s2Count`, it indicates that a permutation of `s1` is present as a substring in `s2`.
