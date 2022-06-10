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

使用滑动窗口搜索：

- 若没有重复的字符，则扩大右边界；
- 若出现重复的字符，则缩小左边界；
- 每次移动时，计算当前子字符串长度，并尝试更新最大子字符串长度。

**References:**

- https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0003.Longest-Substring-Without-Repeating-Characters/
- https://leetcode.com/problems/longest-substring-without-repeating-characters/discuss/376363/CPP-Solution-for-beginners-or-O(n)-time-or-Longest-Substring-without-repeating-characters
