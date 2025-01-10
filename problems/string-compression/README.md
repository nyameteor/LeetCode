# 443. String Compression

- Difficulty: Medium
- Topics: Two Pointers, String
- Link: https://leetcode.com/problems/string-compression/

## Description

Given an array of characters `chars`, compress it using the following algorithm:

Begin with an empty string `s`. For each group of **consecutive repeating characters** in `chars`:

- If the group's length is `1`, append the character to `s`.
- Otherwise, append the character followed by the group's length.

The compressed string `s` **should not be returned separately**, but instead, be stored **in the input character array `chars`**. Note that group lengths that are `10` or longer will be split into multiple characters in `chars`.

After you are done **modifying the input array,** return *the new length of the array*.

You must write an algorithm that uses only constant extra space.

**Example 1:**

```
**Input:** chars = ["a","a","b","b","c","c","c"]
**Output:** Return 6, and the first 6 characters of the input array should be: ["a","2","b","2","c","3"]
**Explanation:** The groups are "aa", "bb", and "ccc". This compresses to "a2b2c3".

```

**Example 2:**

```
**Input:** chars = ["a"]
**Output:** Return 1, and the first character of the input array should be: ["a"]
**Explanation:** The only group is "a", which remains uncompressed since it's a single character.

```

**Example 3:**

```
**Input:** chars = ["a","b","b","b","b","b","b","b","b","b","b","b","b"]
**Output:** Return 4, and the first 4 characters of the input array should be: ["a","b","1","2"].
**Explanation:** The groups are "a" and "bbbbbbbbbbbb". This compresses to "ab12".
```

**Constraints:**

- `1 <= chars.length <= 2000`
- `chars[i]` is a lowercase English letter, uppercase English letter, digit, or symbol.

## Solution

This problem involves compressing a character array using [**run-length encoding**](https://en.wikipedia.org/wiki/Run-length_encoding) in-place, with constant extra space. The solution can be broken down into the following steps:

### Approach: Two Pointers

#### 1. Two Pointers

- **`readIndex`**: Used to traverse the input array.
- **`writeIndex`**: Tracks the position where compressed characters and counts are written.

#### 2. Count Groups

For each group of consecutive characters:

- **Count occurrences**: Use a loop to count how many times the current character repeats.
- **Write character**: Write the character at the current `writeIndex`.
- **Write count (if applicable)**:
  - If the count is greater than 1, convert it to a string.
  - Write each digit of the count to the array at `writeIndex`.

#### 3. Return Compressed Length

After processing all characters, the value of `writeIndex` represents the length of the compressed array.
