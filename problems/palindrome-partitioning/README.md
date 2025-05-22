# 131. Palindrome Partitioning

- Difficulty: Medium
- Topics: String, Dynamic Programming, Backtracking
- Link: https://leetcode.com/problems/palindrome-partitioning/

## Description

Given a string `s`, partition `s` such that every substring of the partition is a **palindrome**. Return *all possible palindrome partitioning of* `s`.

**Example 1:**

```
Input: s = "aab"
Output: [["a","a","b"],["aa","b"]]
```

**Example 2:**

```
Input: s = "a"
Output: [["a"]]
```

**Constraints:**

- `1 <= s.length <= 16`
- `s` contains only lowercase English letters.

## Solution

### Key Idea

- Use backtracking (DFS) to generate all possible partitions.
- Check each substring for palindrome before continuing.
- **Optimization**: memoize palindrome checks to avoid repeated work.
