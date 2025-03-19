# 22. Generate Parentheses

- Difficulty: Medium
- Topics: String, Dynamic Programming, Backtracking
- Link: https://leetcode.com/problems/generate-parentheses/

## Description

Given `n` pairs of parentheses, write a function to _generate all combinations of well-formed parentheses_.

**Example 1:**

```
Input: n = 3
Output: ["((()))","(()())","(())()","()(())","()()()"]
```

**Example 2:**

```
Input: n = 1
Output: ["()"]
```

**Constraints:**

- `1 <= n <= 8`

## Solution

### Approach: Backtracking

1. Use **backtracking** with two counters (`open` and `close`) to track parentheses.
2. Add `'('` if `open < n`, and `')'` if `close < open`.
3. When both counters reach `n`, add the path to the result.
4. **Backtrack** by removing the last parenthesis after each recursive call.

This ensures that only valid parentheses combinations are generated.
