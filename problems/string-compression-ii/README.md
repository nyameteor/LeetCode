# 1531. String Compression II

- Difficulty: Hard
- Topics: String, Dynamic Programming
- Link: https://leetcode.com/problems/string-compression-ii/

## Description

[Run-length encoding](http://en.wikipedia.org/wiki/Run-length_encoding)
 is a string compression method that works by replacing consecutive
identical characters (repeated 2 or more times) with the concatenation
of the character and the number marking the count of the characters
(length of the run). For example, to compress the string `"aabccc"` we replace `"aa"` by `"a2"` and replace `"ccc"` by `"c3"`. Thus the compressed string becomes `"a2bc3"`.

Notice that in this problem, we are not adding `'1'` after single characters.

Given a string `s` and an integer `k`. You need to delete **at most** `k` characters from `s` such that the run-length encoded version of `s` has minimum length.

Find the *minimum length of the run-length encoded version of* `s` *after deleting at most* `k` *characters*.

**Example 1:**

```
**Input:** s = "aaabcccd", k = 2
**Output:** 4
**Explanation:** Compressing s without deleting anything will give us "a3bc3d" of length 6. Deleting any of the characters 'a' or 'c' would at most decrease the length of the compressed string to 5, for instance delete 2 'a' then we will have s = "abcccd" which compressed is abc3d. Therefore, the optimal way is to delete 'b' and 'd', then the compressed version of s will be "a3c3" of length 4.
```

**Example 2:**

```
**Input:** s = "aabbaa", k = 2
**Output:** 2
**Explanation:** If we delete both 'b' characters, the resulting compressed string would be "a4" of length 2.

```

**Example 3:**

```
**Input:** s = "aaaaaaaaaaa", k = 0
**Output:** 3
**Explanation:** Since k is zero, we cannot delete anything. The compressed string is "a11" of length 3.

```

**Constraints:**

- `1 <= s.length <= 100`
- `0 <= k <= s.length`
- `s` contains only lowercase English letters.

## Solution

This problem involves minimizing the length of the run-length encoded string after deleting at most `k` characters from the input string `s`. The solution employs **dynamic programming (DP)** and **greedy strategies** to optimize the compression by considering deletions and merging consecutive character groups.

### Key Concepts

1. **Run-Length Encoding (RLE):**
   - The objective is to compress the string by replacing consecutive identical characters with the character followed by its frequency (e.g., "aaab" becomes "a3b").
   - We are allowed to delete up to `k` characters to minimize the length of the compressed string.

2. **Dynamic Programming Subproblem:**
   - Define `dp(i, k)` as the minimum length of the compressed string for the suffix starting from index `i` after deleting at most `k` characters.
   - For each prefix `s[0..i]`, compute the minimum compressed length by considering the required deletions.

3. **Greedy Group Selection:**
   - For the substring `s[0..i]`, group consecutive characters together and calculate how many deletions are needed to form a valid run-length encoded group.
   - The greedy strategy is to choose the character with the maximum frequency within the current group. This minimizes deletions and maximizes the deletions available for the remaining part of the string.

4. **Recursive Transition:**
   - For each possible partition of the string into two parts (`s[0..i]` and `s[i+1..n)`), calculate the optimal compression by considering deletions in the prefix and recursively solving for the suffix.

5. **Memoization:**
   - To avoid redundant computations, use a memoization table `memo[i][k]` to store the results of subproblems, where `i` is the start index of the suffix and `k` is the number of allowed deletions.

### References

- https://leetcode.com/problems/string-compression-ii/solutions/756022/c-top-down-dp-with-explanation-64ms-short-and-clear
- https://leetcode.com/problems/string-compression-ii/solutions/756022/c-top-down-dp-with-explanation-64ms-short-and-clear/comments/1850831
