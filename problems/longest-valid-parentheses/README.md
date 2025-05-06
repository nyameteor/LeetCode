# 32. Longest Valid Parentheses

- Difficulty: Hard
- Topics: String, Dynamic Programming, Stack
- Link: https://leetcode.com/problems/longest-valid-parentheses/

## Description

Given a string containing just the characters `'('` and `')'`, find the length of the longest valid (well-formed) parentheses substring.

**Example 1:**

```
Input: s = "(()"
Output: 2
Explanation: The longest valid parentheses substring is "()".
```

**Example 2:**

```
Input: s = ")()())"
Output: 4
Explanation: The longest valid parentheses substring is "()()".
```

**Example 3:**

```
Input: s = ""
Output: 0
```

**Constraints:**

- `0 <= s.length <= 3 * 10^4`
- `s[i]` is `'('`, or `')'`.

## Solution

### Stack-Based Approach

We can use a stack to track indices and compute lengths of valid parentheses substrings.

**Key ideas:**

- Push `-1` initially to handle edge cases.
- For `'('`: push its index.
- For `')'`: pop the stack.
  - If the stack is empty, push current index (marks new base).
  - Else, update max length as `current index - stack.top()`.

This allows us to both validate the sequence and compute the longest valid substring in one pass.

**Example walkthrough:**

```text
Input: ")())((())"

Stack:  -1
Step 1: ')' -> pop -> stack empty -> push 0       // Stack: [0]
Step 2: '(' -> push 1                             // Stack: [0, 1]
Step 3: ')' -> pop -> maxLen = 2 - 0 = 2          // Stack: [0]
Step 4: ')' -> pop -> stack empty -> push 3       // Stack: [3]
Step 5: '(' -> push 4                             // Stack: [3, 4]
Step 6: '(' -> push 5                             // Stack: [3, 4, 5]
Step 7: '(' -> push 6                             // Stack: [3, 4, 5, 6]
Step 8: ')' -> pop -> maxLen = 7 - 5 = 2          // Stack: [3, 4, 5]
Step 9: ')' -> pop -> maxLen = 8 - 4 = 4          // Stack: [3, 4]

Final max length = 4
```

### Dynamic Programming

We can use a `dp` array, where `dp[i]` stores the length of the longest valid parentheses substring ending at index `i`.

**Key ideas:**

- If `s[i] == ')'`:
  - Case 1: If `s[i-1] == '('`, then it's a simple `"()"` pair -> `dp[i] = dp[i-2] + 2`
  - Case 2: If `s[i-1] == ')'` and the character **before** the previous valid substring is `'('` (`s[i - dp[i-1] - 1] == '('`), then it wraps around a valid block -> `dp[i] = dp[i-1] + 2 + dp[i - dp[i-1] - 2]`

Iterate through the string, update `dp[i]` accordingly, and track the maximum value.

### References

- [Longest Valid Parentheses | Short & Easy w/ Explanation using stack](https://leetcode.com/problems/longest-valid-parentheses/solutions/1139990/longest-valid-parentheses-short-easy-w-explanation-using-stack/)
