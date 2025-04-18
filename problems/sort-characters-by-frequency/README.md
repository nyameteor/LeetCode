# 451. Sort Characters By Frequency

- Difficulty: Medium
- Topics: Hash Table, String, Sorting, Heap(Priority Queue), Bucket Sort, Counting
- Link: https://leetcode.com/problems/sort-characters-by-frequency/

## Description

Given a string `s`, sort it in **decreasing order** based on the **frequency** of the characters. The **frequency** of a character is the number of times it appears in the string.

Return _the sorted string_. If there are multiple answers, return _any of them_.

**Example 1:**

```
Input: s = "tree"
Output: "eert"
Explanation: 'e' appears twice while 'r' and 't' both appear once.
So 'e' must appear before both 'r' and 't'. Therefore "eetr" is also a valid answer.
```

**Example 2:**

```
Input: s = "cccaaa"
Output: "aaaccc"
Explanation: Both 'c' and 'a' appear three times, so both "cccaaa" and "aaaccc" are valid answers.
Note that "cacaca" is incorrect, as the same characters must be together.
```

**Example 3:**

```
Input: s = "Aabb"
Output: "bbAa"
Explanation: "bbaA" is also a valid answer, but "Aabb" is incorrect.
Note that 'A' and 'a' are treated as two different characters.
```

**Constraints:**

- `1 <= s.length <= 5 * 10^5`
- `s` consists of uppercase and lowercase English letters and digits.

## Solution

- **Count frequencies** of all characters using a hash map.
- **Bucket Sort**:
  - Create buckets where index `i` holds characters that appear `i` times.
  - Iterate buckets in reverse (high to low frequency) and build the result string.
  - Complexity: `O(n)` time, `O(n)` space.
- **Alternative**: Use sorting with custom comparator or a max heap (`O(n log k)` time and `O(n)` space, where `k` is the number of unique characters).
- All characters with the same frequency must be grouped together.
