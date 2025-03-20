# 516. Longest Palindromic Subsequence

- Difficulty: Medium
- Topics: String, Dynamic Programming
- Link: https://leetcode.com/problems/longest-palindromic-subsequence/description/

## Description

Given a string `s`, find *the longest palindromic **subsequence**'s length in* `s`.

A **subsequence** is a sequence that can be derived from
another sequence by deleting some or no elements without changing the
order of the remaining elements.

**Example 1:**

```
Input: s = "bbbab"
Output: 4
Explanation: One possible longest palindromic subsequence is "bbbb".
```

**Example 2:**

```
Input: s = "cbbd"
Output: 2
Explanation: One possible longest palindromic subsequence is "bb".
```

**Constraints:**

- `1 <= s.length <= 1000`
- `s` consists only of lowercase English letters.

## Solution

The problem requires finding the longest palindromic **subsequence**, meaning we can skip characters but must maintain order.

### Intuition

- If the first and last characters match, they contribute to the subsequence:  
   `dp(l, r) = 2 + dp(l+1, r-1)`
- If they don’t match, we must choose:  
   `dp(l, r) = max(dp(l+1, r), dp(l, r-1))`
- Base cases:
  - A single character (`l == r`) is a palindrome of length `1`.
  - An empty range (`l > r`) has length `0`.

### Approach: Dynamic Programming

- Use **top-down DP (memoization)** to store results of `(l, r)` pairs and avoid redundant calculations.
- The recurrence ensures an **O(n²) time complexity** with **O(n²) space** for the memo table.
