# 187. Repeated DNA Sequences

- Difficulty: Medium
- Topics: Hash Table, String, Bit Manipulation, Sliding Window, Rolling Hash, Hash Function
- Link: https://leetcode.com/problems/repeated-dna-sequences/

## Description

The **DNA sequence** is composed of a series of nucleotides abbreviated as `'A'`, `'C'`, `'G'`, and `'T'`.

- For example, `"ACGAATTCCG"` is a **DNA sequence**.

When studying **DNA**, it is useful to identify repeated sequences within the DNA.

Given a string `s` that represents a **DNA sequence**, return all the **`10`-letter-long** sequences (substrings) that occur more than once in a DNA molecule. You may return the answer in **any order**.

**Example 1:**

```
Input: s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
Output: ["AAAAACCCCC","CCCCCAAAAA"]
```

**Example 2:**

```
Input: s = "AAAAAAAAAAAAA"
Output: ["AAAAAAAAAA"]
```

**Constraints:**

- `1 <= s.length <= 10^5`
- `s[i]` is either `'A'`, `'C'`, `'G'`, or `'T'`.

## Solution

### Brute Force

思路：使用 Hash Table 统计每个可能序列的出现次数，如果一个序列出现不止一次，则将其添加到结果中。
