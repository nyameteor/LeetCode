# 140. Word Break II

- Difficulty: Hard
- Topics: Array, Hash Table, String, Dynamic Programming, Backtracking, Trie, Memoization
- Link: https://leetcode.com/problems/word-break-ii/

## Description

Given a string `s` and a dictionary of strings `wordDict`, add spaces in `s` to construct a sentence where each word is a valid dictionary word. Return all such possible sentences in **any order**.

**Note** that the same word in the dictionary may be reused multiple times in the segmentation.

**Example 1:**

```
Input: s = "catsanddog", wordDict = ["cat","cats","and","sand","dog"]
Output: ["cats and dog","cat sand dog"]
```

**Example 2:**

```
Input: s = "pineapplepenapple", wordDict = ["apple","pen","applepen","pine","pineapple"]
Output: ["pine apple pen apple","pineapple pen apple","pine applepen apple"]
Explanation: Note that you are allowed to reuse a dictionary word.
```

**Example 3:**

```
Input: s = "catsandog", wordDict = ["cats","dog","sand","and","cat"]
Output: []
```

**Constraints:**

- `1 <= s.length <= 20`
- `1 <= wordDict.length <= 1000`
- `1 <= wordDict[i].length <= 10`
- `s` and `wordDict[i]` consist of only lowercase English letters.
- All the strings of `wordDict` are **unique**.
- Input is generated in a way that the length of the answer doesn't exceed 10^5.

## Solution

### Approach: Recursive with Memoization

1. **Recursive Decomposition**:
   - Try to match each prefix of `s` with a word in `wordDict`.
   - If a match is found, recursively process the remaining suffix.

2. **Memoization for Optimization**:
   - Store computed results for each substring `s[i:]` in a map (`memo`).
   - If a substring `s[i:]` is already computed, return the stored results instead of recalculating.

3. **Base Case**:
   - When `s` is empty, return `[""]`, indicating that a valid segmentation has been formed.

4. **Sentence Formation**:
   - Construct sentences by concatenating words with spaces.

#### **Complexity Analysis**

- **Time Complexity**: **O(N * 2^N)** in the worst case, where `N` is the length of `s`.
  - The exponential factor arises from the number of possible partitions.
  - Memoization helps reduce redundant recomputation.
- **Space Complexity**: **O(N * 2^N)** due to the storage of valid sentence combinations.
