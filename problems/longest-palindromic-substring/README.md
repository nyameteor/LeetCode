# 5. Longest Palindromic Substring

- Difficulty: Medium
- Topics: String, Dynamic Programming
- Link: https://leetcode.com/problems/longest-palindromic-substring/

## Description

Given a string `s`, return _the longest palindromic substring_ in `s`.

**Example 1:**

```
Input: s = "babad"
Output: "bab"
Explanation: "aba" is also a valid answer.
```

**Example 2:**

```
Input: s = "cbbd"
Output: "bb"
```

**Constraints:**

- `1 <= s.length <= 1000`
- `s` consist of only digits and English letters.

## Solution

### Two pointers

For each character `s[i]`:

- Assume palindrome has odd length, extend it as much as possible starting from `left = i` and `right = i`.
- Assume palindrome has even length, extend it as much as possible starting from `left = i` and `right = i+1`.

And update the longest palindrome if needed.

References:

- https://leetcode.com/problems/longest-palindromic-substring/solutions/2928/very-simple-clean-java-solution/
- https://leetcode.com/problems/longest-palindromic-substring/solutions/151144/bottom-up-dp-two-pointers/
