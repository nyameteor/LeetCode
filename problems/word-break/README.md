# 139. Word Break

- Difficulty: Medium
- Topics: Hash Table, String, Dynamic Programming, Trie, Memoization
- Link: https://leetcode.com/problems/word-break/

## Description

- 给定一个字符串 s 和一个字符串字典 wordDict，如果 s 可以被分割成一个或多个字典单词的“空格”序列，则返回 true。
- 拆分时可以重复使用字典中的单词。

```shell
Example 1:

Input: s = "leetcode", wordDict = ["leet","code"]
Output: true
Explanation: Return true because "leetcode" can be segmented as "leet code".

Example 2:

Input: s = "applepenapple", wordDict = ["apple","pen"]
Output: true
Explanation: Return true because "applepenapple" can be segmented as "apple pen apple".
Note that you are allowed to reuse a dictionary word.

Example 3:

Input: s = "catsandog", wordDict = ["cats","dog","sand","and","cat"]
Output: false
```

## Solution

字典中的单词可以重复使用，故是一个[无界背包问题](https://zh.wikipedia.org/wiki/%E8%83%8C%E5%8C%85%E9%97%AE%E9%A2%98)。

### Recursion + Memoization

```shell
# 基本情况
s = "" -> f(s, dict) = true

# 递归方程
for x in dict:
  f(s - x, dict) = true -> f(s, dict) = true

# 添加 memo 优化为伪多项式时间
```

- m = s.length, n = dict.length
- time: O(m \* n), Space: O(m)
