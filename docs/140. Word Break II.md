# 140. Word Break II

- Difficulty: Hard
- Topics: Hash Table, String, Dynamic Programming, Trie, Memoization
- Link: https://leetcode.com/problems/word-break-ii/

## Description

- 给定一个字符串 s 和一个字符串字典 wordDict，在 s 中添加空格以构造一个句子，其中每个单词都是一个有效的字典单词。 返回所有这些可能的句子（任意顺序）。
- 注意：词典中的同一个词可能会在切分中多次重复使用。

```shell
Example 1:

Input: s = "catsanddog", wordDict = ["cat","cats","and","sand","dog"]
Output: ["cats and dog","cat sand dog"]

Example 2:

Input: s = "pineapplepenapple", wordDict = ["apple","pen","applepen","pine","pineapple"]
Output: ["pine apple pen apple","pineapple pen apple","pine applepen apple"]
Explanation: Note that you are allowed to reuse a dictionary word.

Example 3:

Input: s = "catsandog", wordDict = ["cats","dog","sand","and","cat"]
Output: []
```

## Solution

和 139 题 "单词拆分" 相似，139 属于 `Decision Problem` 返回判断结果；该题属于 `Combinatoric Problem` 求所有组合，思路基本相同。

### Recursion + Memoization

```shell
# 基本情况
s = "" -> f(s, dict) = [""]

# 递归方程
for x in dict:
  f(s - x, dict) != [] -> f(s, dict) = f(s - x, dict).map(e -> e + x)

# 遍历累积
# res 累积 f(s, dict) 递归遍历中的所有结果，遍历结束后返回 res

# 添加 memo 优化为伪多项式时间
```
