# 139. Word Break

- Difficulty: Medium
- Topics: Hash Table, String, Dynamic Programming, Trie, Memoization
- Link: https://leetcode.com/problems/word-break/

## Description

Given a string `s` and a dictionary of strings `wordDict`, return `true` if `s` can be segmented into a space-separated sequence of one or more dictionary words.

**Note** that the same word in the dictionary may be reused multiple times in the segmentation.

**Example 1:**

```
Input: s = "leetcode", wordDict = ["leet","code"]
Output: true
Explanation: Return true because "leetcode" can be segmented as "leet code".
```

**Example 2:**

```
Input: s = "applepenapple", wordDict = ["apple","pen"]
Output: true
Explanation: Return true because "applepenapple" can be segmented as "apple pen apple".
Note that you are allowed to reuse a dictionary word.
```

**Example 3:**

```
Input: s = "catsandog", wordDict = ["cats","dog","sand","and","cat"]
Output: false
```

**Constraints:**

- `1 <= s.length <= 300`
- `1 <= wordDict.length <= 1000`
- `1 <= wordDict[i].length <= 20`
- `s` and `wordDict[i]` consist of only lowercase English letters.
- All the strings of `wordDict` are **unique**.

## Solution

### Approach: Recursive with Memoization

We recursively check if `s` can be broken down into words from `wordDict`. At each step, we check if `s` starts with a word from `wordDict`. If so, we recurse on the remaining substring. We use a **memoization map** to store results for substrings to avoid redundant computations.

- **Base Case:** If `s` is empty, return `true`.
- **Memoization:** If `s` has been checked before, return the stored result.
- **Recursive Step:** Try each word in `wordDict`. If `s` starts with `word`, check the remainder. If any split leads to `true`, store and return `true`. Otherwise, return `false`.

**Time Complexity:** ~O(N * M) in the worst case (N = length of `s`, M = wordDict size).

For optimized solutions and performance improvements, please see the [official editorial](https://leetcode.com/problems/word-break/editorial/).
