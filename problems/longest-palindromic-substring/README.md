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

### Dynamic Programming

定义 `P(i, j)` 为子串 `s[i...j]` 是否为回文串。

容易得出递推式为：

```shell
# 若当前子串左右结尾处的两个字符相同，且去掉左右结尾的字符后的子串为回文串，则当前子串也为回文串
P(i, j) = P(i+1, j-1) && s[i] == s[j]
```

由于当前状态与 `P(i+1, j-1)` 有关，仅当 `size(s) >= 3` 时 `i+1` 和 `j-1` 才合法，故存在两种边界情况：

```shell
# 只有一个字符的子串
P(i, i) = true
# 只有两个字符的子串
P(i, i+1) = (s[i] == s[i+1])
```

最后，由于转换函数为 `P(i, j) = P(i+1, j-1)`，由于 i+1 在 i 之前先被求值，j-1 在 j 之前先被求值，因此最好让 i 以递减序列遍历，j 以递增序列遍历。
