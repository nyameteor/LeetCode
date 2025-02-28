# 1092. Shortest Common Supersequence

- Difficulty: Hard
- Topics: String, Dynamic Programming
- Link: https://leetcode.com/problems/shortest-common-supersequence/

## Description

Given two strings `str1` and `str2`, return *the shortest string that has both* `str1` *and* `str2` *as **subsequences***. If there are multiple valid strings, return **any** of them.

A string `s` is a **subsequence** of string `t` if deleting some number of characters from `t` (possibly `0`) results in the string `s`.

**Example 1:**

```
Input: str1 = "abac", str2 = "cab"
Output: "cabac"
Explanation: 
str1 = "abac" is a subsequence of "cabac" because we can delete the first "c".
str2 = "cab" is a subsequence of "cabac" because we can delete the last "ac".
The answer provided is the shortest such string that satisfies these properties.

```

**Example 2:**

```
Input: str1 = "aaaaaaaa", str2 = "aaaaaaaa"
Output: "aaaaaaaa"

```

**Constraints:**

- `1 <= str1.length, str2.length <= 1000`
- `str1` and `str2` consist of lowercase English letters.

## Solution

This solution finds the **Shortest Common Supersequence (SCS)** by leveraging the **Longest Common Subsequence (LCS)** as a guide.

### Steps

1. **Compute LCS using recursive DP with memoization:**  
   - If characters match, include them in LCS and move diagonally (`i+1, j+1`).  
   - Otherwise, take the maximum LCS by moving either `i+1` or `j+1`.
   - Memoization helps avoid redundant calculations.

2. **Build SCS using LCS information:**  
   - Recursively construct the shortest string that includes both `str1` and `str2` as subsequences.  
   - If characters match, append and move diagonally.  
   - Otherwise, append the character from the path leading to the larger LCS.
