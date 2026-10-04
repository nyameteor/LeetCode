# 49. Group Anagrams

- Difficulty: Medium
- Topics: Array, Hash Table, String, Sorting
- Link: https://leetcode.com/problems/group-anagrams/

## Description

Given an array of strings `strs`, group the anagrams together. You can return the answer in **any order**.

**Example 1:**

> **Input:** strs = \["eat","tea","tan","ate","nat","bat"\]
> 
> **Output:** \[\["bat"\],\["nat","tan"\],\["ate","eat","tea"\]\]
> 
> **Explanation:**
> 
> - There is no string in strs that can be rearranged to form `"bat"`.
> - The strings `"nat"` and `"tan"` are anagrams as they can be rearranged to form each other.
> - The strings `"ate"`, `"eat"`, and `"tea"` are anagrams as they can be rearranged to form each other.

**Example 2:**

> **Input:** strs = \[""\]
> 
> **Output:** \[\[""\]\]

**Example 3:**

> **Input:** strs = \["a"\]
> 
> **Output:** \[\["a"\]\]

**Constraints:**

- `1 <= strs.length <= 10^4`
- `0 <= strs[i].length <= 100`
- `strs[i]` consists of lowercase English letters.

## Solution

### Approach: Frequency Count

Count the 26 lowercase letters in each string. Anagrams share that count, so use it as the hash map key and append the string to its group.

#### Complexity

- **Time:** `O(nk)`, where `n` is the number of strings and `k` is the longest length
- **Space:** `O(nk)`
