# 3. Longest Substring Without Repeating Characters

- Difficulty: Medium
- Topics: Hash Table, String, Sliding Window
- Link: https://leetcode.com/problems/longest-substring-without-repeating-characters/

## Description

Given a string `s`, find the length of the **longest substring** without repeating characters.

**Example 1:**

```
Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
```

**Example 2:**

```
Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
```

**Example 3:**

```
Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
```

**Constraints:**

- `0 <= s.length <= 5 * 10^4`
- `s` consists of English letters, digits, symbols and spaces.

## Solution

### Sliding Window

Search using the sliding window:

- If there are no repeated characters, expand the right border;
- If repeated characters appear, reduce the left border;
- On each move, calculate the current substring length and try to update the maximum substring length.

**References:**

- https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0003.Longest-Substring-Without-Repeating-Characters/
- https://leetcode.com/problems/longest-substring-without-repeating-characters/discuss/376363/CPP-Solution-for-beginners-or-O(n)-time-or-Longest-Substring-without-repeating-characters
