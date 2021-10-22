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

### Priority Queue

1. 使用 map `freq` 计数每个字符出现的次数；
2. 使用 priority queue `q` 辅助对 `freq` 进行排序；
3. pop `q` 中的所有元素，构造结果 string 即可。

推荐使用 unorderd_map 而不是 map，因为 unorderd_map 的实现是 hash table，而 map 的实现是 balanced binary tree，前者搜索的时间复杂度更低。

参考：https://hackingcpp.com/cpp/std/associative_containers.html

### Bucket Sort

Todo
