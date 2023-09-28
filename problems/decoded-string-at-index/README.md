# 880. Decoded String at Index

- Difficulty: Medium
- Topics: String, Stack
- Link: https://leetcode.com/problems/decoded-string-at-index/

## Description

You are given an encoded string `s`. To decode the string to a tape, the encoded string is read one character at a time and the following steps are taken:

- If the character read is a letter, that letter is written onto the tape.
- If the character read is a digit `d`, the entire current tape is repeatedly written `d - 1` more times in total.

Given an integer `k`, return _the_ `kth` _letter (**1-indexed)** in the decoded string_.

**Example 1:**

```
**Input:** s = "leet2code3", k = 10
**Output:** "o"
**Explanation:** The decoded string is "leetleetcodeleetleetcodeleetleetcode".
The 10th letter in the string is "o".
```

**Example 2:**

```
**Input:** s = "ha22", k = 5
**Output:** "h"
**Explanation:** The decoded string is "hahahaha".
The 5th letter is "h".
```

**Example 3:**

```
**Input:** s = "a2345678999999999999999", k = 1
**Output:** "a"
**Explanation:** The decoded string is "a" repeated 8301530446056247680 times.
The 1st letter is "a".
```

**Constraints:**

- `2 <= s.length <= 100`
- `s` consists of lowercase English letters and digits `2` through `9`.
- `s` starts with a letter.
- `1 <= k <= 109`
- It is guaranteed that `k` is less than or equal to the length of the decoded string.
- The decoded string is guaranteed to have less than `2^63` letters.

## Solution

Reference: https://leetcode.com/problems/decoded-string-at-index/solutions/4094710/100-reverse-stack-commented-code/

We can solve this problem without really decoding it:

1. Traverse the whole string and get the total `length` of decoded string.
2. Reverse traversing the string and apply two operations, divide the `length` and mod the `k` until we get `k`.

For example, this is the simulation with `s = "leet2code3"`, `k = 26` during reverse traversing:

| char | length | k   | next_length | next_k          |
| ---- | ------ | --- | ----------- | --------------- |
| 3    | 36     | 26  | length / 3  | k % next_length |
| e    | 12     | 2   | length - 1  |                 |
| d    | 11     | 2   | length - 1  |                 |
| o    | 10     | 2   | length - 1  |                 |
| c    | 9      | 2   | length - 1  |                 |
| 2    | 8      | 2   | length / 2  | k % next_length |
| t    | 4      | 2   | length - 1  |                 |
| e    | 3      | 2   | length - 1  |                 |
| e    | 2      | 2   | length - 1  |                 |

So the result is `e`.
