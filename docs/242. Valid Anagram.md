# 242. Valid Anagram

- Difficulty: Easy
- Topics: Hash Table, String, Sorting
- Link: https://leetcode.com/problems/valid-anagram/

## Description

Given two strings `s` and `t`, return `true` _if_ `t` _is an anagram of_ `s`_, and_ `false` _otherwise_.

**Example 1:**

```
Input: s = "anagram", t = "nagaram"
Output: true
```

**Example 2:**

```
Input: s = "rat", t = "car"
Output: false
```

**Constraints:**

- `1 <= s.length, t.length <= 5 * 10^4`
- `s` and `t` consist of lowercase English letters.

**Follow up:** What if the inputs contain Unicode characters? How would you adapt your solution to such a case?

## Solution

给定两个字符串 s 和 t ，判断 t 是否是 s 的字母异位词。

### Hash Table

若两个字符串是字母异位词，说明两个字符串中所有字符出现的次数相同，字符顺序可以相同也可以不同。可以想到用散列表来存储每个字符出现的次数。

[散列表](https://en.wikipedia.org/wiki/Hash_table) 的定义即是关联数组，和数组不同的是散列表可以自动扩容。因为本题限定了字符串只包含小写字母，所以表的大小是固定的，也可用一般的数组。

使用一个数组统计 s 和 t 中各个字母出现的次数，将 ASCII 相对值 `字符 - 'a'` 作为数组下标，当 s 出现时字符出现次数 +1，t 中出现时字符出现次数 -1，最后遍历数组若全为 0，说明 s 和 t 为字母异位词。
