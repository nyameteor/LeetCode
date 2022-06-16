# 583. Delete Operation for Two Strings

- Difficulty: Medium
- Topics: String, Dynamic Programming
- Link: https://leetcode.com/problems/delete-operation-for-two-strings/

## Description

Given two strings `word1` and `word2`, return _the minimum number of **steps** required to make_ `word1` _and_ `word2` _the same_.

In one **step**, you can delete exactly one character in either string.

**Example 1:**

```
Input: word1 = "sea", word2 = "eat"
Output: 2
Explanation: You need one step to make "sea" to "ea" and another step to make "eat" to "ea".
```

**Example 2:**

```
Input: word1 = "leetcode", word2 = "etco"
Output: 4
```

**Constraints:**

- `1 <= word1.length, word2.length <= 500`
- `word1` and `word2` consist of only lowercase English letters.

## Solution

Denote length of s1 as $len1$, length of s2 as $len2$.

Denote the longest common sequence among the two strings as $lcs$.

We can easily determine that $result = len1 + len2 - 2 * lcs$.

To search $lcs$, we can use a recursive function $lcs(s1, s2, i1, i2)$.

$i1$ and $i2$ is the $lcs$ among the string $s1$ and $s2$ considering their tail indexes upto $i1$ and $i2$ respectively.

$lcs(s1, s2, i1, i2)$ can be evaluated as:

- if $i1 = -1 \text{ or } i2 = -1$, return $0$.
- if $s1[i1] = s2[i2]$, return $1 + lcs(s1, s2, i1-1, i2-1)$.
- if $s1[i1] \ne s2[i2]$, return $max(lcs(s1, s2, i1-1, i2), lcs(s1, s2, i1, i2-1))$.

Then to avoid redundant function calls, we can add $memo$ array to store calculated values.

**References:**

- https://leetcode.com/problems/delete-operation-for-two-strings/solution/
