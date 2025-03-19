# 438. Find All Anagrams in a String

- Difficulty: Medium
- Topics: Hash Table, String, Sliding Window
- Link: https://leetcode.com/problems/find-all-anagrams-in-a-string/

## Description

Given two strings `s` and `p`, return _an array of all the start indices of_ `p`_'s anagrams in_ `s`. You may return the answer in **any order**.

An **Anagram** is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

**Example 1:**

```
Input: s = "cbaebabacd", p = "abc"
Output: [0,6]
Explanation:
The substring with start index = 0 is "cba", which is an anagram of "abc".
The substring with start index = 6 is "bac", which is an anagram of "abc".
```

**Example 2:**

```
Input: s = "abab", p = "ab"
Output: [0,1,2]
Explanation:
The substring with start index = 0 is "ab", which is an anagram of "ab".
The substring with start index = 1 is "ba", which is an anagram of "ab".
The substring with start index = 2 is "ab", which is an anagram of "ab".
```

**Constraints:**

- `1 <= s.length, p.length <= 3 * 10^4`
- `s` and `p` consist of lowercase English letters.

## Solution

### Sliding Window

### Approach: Sliding Window + Frequency Count

This solution efficiently finds anagrams using the **Sliding Window technique** with **constant space (O(1))** for character frequency tracking.

1. **Frequency Count Initialization**
   - Create two frequency arrays (`pFreq`, `windowFreq`) of size `26` (for lowercase English letters).
   - Populate them with the frequency of `p` and the first window in `s`.

2. **Matching Character Count (`matchCount`)**
   - Count how many characters have matching frequencies in `pFreq` and `windowFreq`.

3. **Sliding Window Processing**
   - Move the window one step at a time.
   - **Remove** the outgoing left character and **add** the incoming right character.
   - Update `matchCount` dynamically, ensuring we only compare affected characters.

4. **Efficiently Check for Anagram Match**
   - If `matchCount == 26`, the window is a valid anagram.
