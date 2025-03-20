# 5. Longest Palindromic Substring

- Difficulty: Medium
- Topics: String, Dynamic Programming
- Link: https://leetcode.com/problems/longest-palindromic-substring/

## Description

Given a string `s`, return _the longest palindromic substring_ in `s`.

**Example 1:**

```
Input: s = "babad"
Output: "bab"
Explanation: "aba" is also a valid answer.
```

**Example 2:**

```
Input: s = "cbbd"
Output: "bb"
```

**Constraints:**

- `1 <= s.length <= 1000`
- `s` consist of only digits and English letters.

## Solution

### Approach: Two pointers (Expand-Around-Center)

For each character `s[i]`:

- Assume palindrome has **odd** length, extend it as much as possible starting from `left = i` and `right = i`.
- Assume palindrome has **even** length, extend it as much as possible starting from `left = i` and `right = i+1`.
- Update the longest palindrome found.

**Complexity**:

- Time: `O(n^2)`.
- Space: `O(1)`.

### Approach: Dynamic Programming (Bottom-Up)

- Use `dp[i][j]` to track if `s[i:j+1]` is a palindrome.
- Single characters are palindromes (`dp[i][i] = true`).
- Two-character substrings are palindromes if both match.
- Longer substrings are palindromes if `s[l] == s[r]` and `dp[l+1][r-1]` is `true`.
- Track the longest palindrome found.

**Complexity**:

- Time: `O(n^2)`, iterating over all substrings.
- Space: `O(n^2)`, storing palindrome states.

**Note**: Slower than Expand-Around-Center due to extra space but useful for educational purposes.

### References

- https://leetcode.com/problems/longest-palindromic-substring/solutions/2928/very-simple-clean-java-solution/
- https://leetcode.com/problems/longest-palindromic-substring/solutions/151144/bottom-up-dp-two-pointers/
